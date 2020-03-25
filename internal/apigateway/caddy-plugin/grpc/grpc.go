package grpc

import (
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"hal9000/internal/apigateway/caddy-plugin/internal"
	"net/http"
)

type RpcProxy struct {
	Rule *Rule
	Next httpserver.Handler
}

type Rule struct {
	Path           string
	URL            string
	Handler        http.Handler
	ExclusionRules []internal.ExclusionRule
}

func (r RpcProxy) ServeHTTP(resp http.ResponseWriter, req *http.Request) (int, error) {
	for _, rule := range r.Rule.ExclusionRules {
		if httpserver.Path(req.URL.Path).Matches(rule.Path) && (rule.Method == internal.AllMethod || req.Method == rule.Method) {
			return r.Next.ServeHTTP(resp, req)
		}
	}

	if httpserver.Path(req.URL.Path).Matches(r.Rule.URL) {
		r.Rule.Handler.ServeHTTP(resp, req)
		return http.StatusOK, nil
	}

	return r.Next.ServeHTTP(resp, req)
}
