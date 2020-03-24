package grpc

import (
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"net/http"
)

type RpcProxy struct {
	Rule *Rule
	Next httpserver.Handler
}


type Rule struct {
	Path                string
}

func (r RpcProxy) ServeHTTP(resp http.ResponseWriter, req *http.Request) (int, error) {
	return 0, nil
}