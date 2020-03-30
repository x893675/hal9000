package app

import (
	"github.com/spf13/cobra"
	"hal9000/cmd/account/app/options"
	"hal9000/internal/account"
	"hal9000/pkg/client"
	serverconfig "hal9000/pkg/httpserver/config"
	"hal9000/pkg/logger"
	"hal9000/pkg/utils/signals"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func NewAccountServiceCommand() *cobra.Command {
	s := options.NewAccountServiceOptions()

	cmd := &cobra.Command{
		Use:  "account-service",
		Long: `account service`,
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
	fs.StringVar(&s.Loglevel, "loglevel", s.Loglevel, "info server log level, e.g. debug,info")
	namedFlagSets := s.Flags()

	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return cmd
}

// apply server run options to configuration
func Complete(s *options.AccountServiceOptions) error {

	// loading configuration file
	conf := serverconfig.Get()

	conf.Apply(&serverconfig.Config{
		DatabaseOptions: s.DatabaseOptions,
	})

	*s = options.AccountServiceOptions{
		DatabaseOptions: conf.DatabaseOptions,
		Loglevel:        s.Loglevel,
	}

	return nil
}

func Run(s *options.AccountServiceOptions, stopCh <-chan struct{}) error {
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

func CreateAccountService(s *options.AccountServiceOptions) error {
	account.Serve()
	return nil
}