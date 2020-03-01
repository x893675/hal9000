package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"hal9000/cmd/api-server/app/options"
	"hal9000/internal/apiserver"
	"hal9000/pkg/client"
	"hal9000/pkg/logger"
	"hal9000/pkg/server"
	serverconfig "hal9000/pkg/server/config"
	"hal9000/pkg/server/filter"
	"hal9000/pkg/server/runtime"
	"hal9000/pkg/server/version"
	"hal9000/pkg/utils/signals"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"net/http"
)

func NewAPIServerCommand() *cobra.Command {
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

	conf.Apply(&serverconfig.Config{
		MySQLOptions: s.MySQLOptions,
	})

	*s = options.ServerRunOptions{
		GenericServerRunOptions: s.GenericServerRunOptions,
		MySQLOptions:            conf.MySQLOptions,
	}

	return nil
}

func Run(s *options.ServerRunOptions, stopCh <-chan struct{}) error {

	err := CreateClientSet(serverconfig.Get(), stopCh)
	if err != nil {
		return err
	}

	err = CreateAPIServer(s)
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

func CreateAPIServer(s *options.ServerRunOptions) error {
	var err error

	container := runtime.Container
	container.DoNotRecover(false)
	container.Filter(filter.Logging)
	container.RecoverHandler(server.LogStackOnRecover)

	apiserver.InstallAPIs(container)

	// install config api
	serverconfig.InstallAPI(container)

	if s.GenericServerRunOptions.InsecurePort != 0 {
		logger.Info(nil, "Server [version: %s] Start on %s:%d", version.Version, s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.InsecurePort)
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.InsecurePort), container)
		if err == nil {
			logger.Info(nil, "Server listening on insecure port %d.", s.GenericServerRunOptions.InsecurePort)
		}
	}

	if s.GenericServerRunOptions.SecurePort != 0 && len(s.GenericServerRunOptions.TlsCertFile) > 0 && len(s.GenericServerRunOptions.TlsPrivateKey) > 0 {
		err = http.ListenAndServeTLS(fmt.Sprintf("%s:%d", s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.SecurePort), s.GenericServerRunOptions.TlsCertFile, s.GenericServerRunOptions.TlsPrivateKey, container)
		if err == nil {
			logger.Info(nil, "Server listening on secure port %d.", s.GenericServerRunOptions.SecurePort)
		}
	}

	return err
}
