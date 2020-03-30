package app

import (
	"github.com/spf13/cobra"
	"hal9000/cmd/auth/app/options"
	"hal9000/internal/auth"
	"hal9000/pkg/client"
	serverconfig "hal9000/pkg/httpserver/config"
	"hal9000/pkg/logger"
	"hal9000/pkg/utils/signals"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func NewAuthServiceCommand() *cobra.Command {
	s := options.NewAuthServiceOptions()

	cmd := &cobra.Command{
		Use:  "auth-service",
		Long: `auth service`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := serverconfig.Load()
			if err != nil {
				return err
			}

			err = Complete(s)
			if err != nil {
				return err
			}

			if errs := s.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}

			return Run(s, signals.SetupSignalHandler())
		},
	}

	fs := cmd.Flags()
	namedFlagSets := s.Flags()

	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return cmd
}

// apply server run options to configuration
func Complete(s *options.AuthServiceOptions) error {

	// loading configuration file
	conf := serverconfig.Get()

	conf.Apply(&serverconfig.Config{
		DatabaseOptions: s.DatabaseOptions,
	})

	*s = options.AuthServiceOptions{
		DatabaseOptions: conf.DatabaseOptions,
		Loglevel:        s.Loglevel,
	}

	return nil
}

func Run(s *options.AuthServiceOptions, stopCh <-chan struct{}) error {
	logger.SetLevelByString(s.Loglevel)
	err := CreateClientSet(serverconfig.Get(), stopCh)
	if err != nil {
		return err
	}

	err = CreateAccountService(s)
	if err != nil {
		return err
	}

	return nil
}

func CreateClientSet(conf *serverconfig.Config, stopCh <-chan struct{}) error {
	csop := &client.ClientSetOptions{}

	csop.SetDatabaseOptions(conf.DatabaseOptions).
		SetLdapOptions(conf.LdapOptions)

	client.NewClientSetFactory(csop, stopCh)

	return nil
}

func CreateAccountService(s *options.AuthServiceOptions) error {
	auth.Serve()
	return nil
}
