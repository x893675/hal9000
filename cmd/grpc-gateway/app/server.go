package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"hal9000/cmd/grpc-gateway/app/options"
	"hal9000/pkg/client"
	"hal9000/pkg/constants"
	serverconfig "hal9000/pkg/httpserver/config"
	"hal9000/pkg/httpserver/version"
	"hal9000/pkg/logger"
	"hal9000/pkg/utils/signals"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func NewGrpcGatewayCommand() *cobra.Command {
	s := options.NewServerRunOptions()

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
func Complete(s *options.ServerRunOptions) error {

	// loading configuration file
	conf := serverconfig.Get()

	conf.Apply(&serverconfig.Config{})

	*s = options.ServerRunOptions{
		GenericServerRunOptions: s.GenericServerRunOptions,
		Loglevel:                s.Loglevel,
	}

	return nil
}

func Run(s *options.ServerRunOptions, stopCh <-chan struct{}) error {
	logger.SetLevelByString(s.Loglevel)
	err := CreateClientSet(serverconfig.Get(), stopCh)
	if err != nil {
		return err
	}

	err = CreateGrpcGateway(s)
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

func CreateGrpcGateway(s *options.ServerRunOptions) error {
	var err error

	if s.GenericServerRunOptions.SecurePort != 0 {
		logger.Critical(nil, "grpc gateway run without tls")
		return fmt.Errorf("grpc gateway run without tls")
	}

	logger.Info(nil, "Grpc gateway [version: %s] Start on %s:%d", version.Version, s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.InsecurePort)
	logger.Info(nil, "Test service %s:%d", constants.TestServiceHost, constants.TestServicePort)


	gw := Server{}

	if err = gw.run(s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.InsecurePort); err != nil {
		logger.Critical(nil, "Grpc Gateway run failed %+v", err)
	}

	return err
}