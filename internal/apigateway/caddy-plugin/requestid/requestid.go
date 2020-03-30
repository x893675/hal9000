package requestid

import (
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"github.com/google/uuid"
	"hal9000/pkg/logger"
	"net/http"
)

type Rule struct {
	key  string
	Next httpserver.Handler
}

func (r Rule) ServeHTTP(resp http.ResponseWriter, req *http.Request) (int, error) {
	logger.Info(nil, "key is [%v]", r.key)
	requestId := req.Header.Get(r.key)
	logger.Info(nil, "key is [%v]", requestId)
	if requestId == "" {
		req.Header.Set(r.key, uuid.New().String())
	}
	logger.Info(nil, "request id is [%v]", req.Header.Get("x-request-id"))
	logger.Info(nil, "request id is [%v]", req.Header.Get("X-Request-ID"))
	return r.Next.ServeHTTP(resp, req)
}
