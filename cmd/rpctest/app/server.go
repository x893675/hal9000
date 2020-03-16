package app

import (
	"github.com/spf13/cobra"
	"hal9000/cmd/rpctest/app/options"
	"hal9000/internal/rpctest"
	"hal9000/pkg/client"
	serverconfig "hal9000/pkg/httpserver/config"
	"hal9000/pkg/logger"
	"hal9000/pkg/utils/signals"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func NewTestServerCommand() *cobra.Command {
	s := options.NewServiceRunOptions()

	cmd := &cobra.Command{
		Use:  "api-server",
		Long: `restful api server`,
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
func Complete(s *options.ServiceRunOptions) error {

	// loading configuration file
	conf := serverconfig.Get()

	conf.Apply(&serverconfig.Config{
		MySQLOptions: s.MySQLOptions,
	})

	*s = options.ServiceRunOptions{
		MySQLOptions:            conf.MySQLOptions,
		Loglevel:                s.Loglevel,
	}

	return nil
}

func Run(s *options.ServiceRunOptions, stopCh <-chan struct{}) error {
	logger.SetLevelByString(s.Loglevel)
	err := CreateClientSet(serverconfig.Get(), stopCh)
	if err != nil {
		return err
	}

	err = CreateTestService(s)
	if err != nil {
		return err
	}

	return nil
}

func CreateClientSet(conf *serverconfig.Config, stopCh <-chan struct{}) error {
	csop := &client.ClientSetOptions{}

	csop.SetMySQLOptions(conf.MySQLOptions).
		SetLdapOptions(conf.LdapOptions)

	client.NewClientSetFactory(csop, stopCh)

	return nil
}

func CreateTestService(s *options.ServiceRunOptions) error {
	rpctest.Serve()
	return nil
}
