package app

import (
	"flag"
	"github.com/spf13/cobra"
	"hal9000/internal/apigateway"
	"hal9000/pkg/client"
	serverconfig "hal9000/pkg/httpserver/config"
	"github.com/caddyserver/caddy/caddy/caddymain"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"hal9000/pkg/utils/signals"
)

func NewAPIGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "api-gateway",
		Long: `The middle platform api gateway.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			err := serverconfig.Load()
			if err != nil {
				return err
			}

			apigateway.RegisterPlugins()

			return Run(signals.SetupSignalHandler())
		},
	}

	cmd.Flags().AddGoFlagSet(flag.CommandLine)

	return cmd
}


func Run(stopCh <-chan struct{}) error {
	csop := &client.ClientSetOptions{}

	client.NewClientSetFactory(csop, stopCh)

	httpserver.RegisterDevDirective("grpcproxy", "jwt")
	//httpserver.RegisterDevDirective("authorize", "jwt")
	//httpserver.RegisterDevDirective("audit", "jwt")
	//httpserver.RegisterDevDirective("swagger", "jwt")

	caddymain.EnableTelemetry = false
	caddymain.Run()

	return nil
}