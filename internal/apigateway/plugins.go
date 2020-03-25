package apigateway

import (
	"github.com/caddyserver/caddy"
	"hal9000/internal/apigateway/caddy-plugin/grpc"
)

func RegisterPlugins() {
	caddy.RegisterPlugin("grpcproxy", caddy.Plugin{
		ServerType: "http",
		Action:     grpc.Setup,
	})
}