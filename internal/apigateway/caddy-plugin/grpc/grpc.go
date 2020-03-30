package grpc

import (
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"hal9000/internal/apigateway/caddy-plugin/internal"
	"hal9000/pkg/logger"
	"net/http"
)

type RpcProxy struct {
	Rule *Rule
	Next httpserver.Handler
}

type Rule struct {
	Path           string
	Handler        http.Handler
	ExclusionRules []internal.ExclusionRule
}

func (r RpcProxy) ServeHTTP(resp http.ResponseWriter, req *http.Request) (int, error) {
	for _, rule := range r.Rule.ExclusionRules {
		if httpserver.Path(req.URL.Path).Matches(rule.Path) && (rule.Method == internal.AllMethod || req.Method == rule.Method) {
			return r.Next.ServeHTTP(resp, req)
		}
	}

	if httpserver.Path(req.URL.Path).Matches(r.Rule.Path) {
		logger.Info(nil, "request url is %v", req.URL.Path)
		logger.Info(nil, "request id is [%v]", req.Header.Get("x-request-id"))
		logger.Info(nil, "request id is [%v]", req.Header.Get("X-Request-ID"))
		r.Rule.Handler.ServeHTTP(resp, req)
		return http.StatusOK, nil
	}

	return r.Next.ServeHTTP(resp, req)
}
