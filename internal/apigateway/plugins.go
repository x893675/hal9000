package apigateway

import (
	"github.com/caddyserver/caddy"
	"hal9000/internal/apigateway/caddy-plugin/grpc"
	"hal9000/internal/apigateway/caddy-plugin/requestid"
)

func RegisterPlugins() {
	caddy.RegisterPlugin("requestid", caddy.Plugin{
		ServerType: "http",
		Action:     requestid.Setup,
	})
	caddy.RegisterPlugin("grpcproxy", caddy.Plugin{
		ServerType: "http",
		Action:     grpc.Setup,
	})
}