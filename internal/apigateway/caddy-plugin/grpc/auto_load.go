package grpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hal9000/internal/apigateway/caddy-plugin/internal"
	"hal9000/pkg/constants"
	"hal9000/pkg/logger"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
	"hal9000/pkg/utils/sliceutils"
	"io/ioutil"
	"net/http"
	"strings"
)

type register struct {
	f        func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	endpoint string
}

func Setup(c *caddy.Controller) error {
	rule, err := parse(c)

	if err != nil {
		return err
	}

	c.OnStartup(func() error {
		fmt.Println("grpc middleware is start...")
		return nil
	})

	c.OnShutdown(func() error {
		fmt.Println("grpc middleware shutdown...")
		return nil
	})

	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return &RpcProxy{Next: next, Rule: rule}
	})

	return nil
}

func parse(c *caddy.Controller) (*Rule, error) {
	rule := &Rule{URL: "/"}

	if c.Next() {
		args := c.RemainingArgs()
		switch len(args) {
		case 0:
			for c.NextBlock() {
				switch c.Val() {
				case "path":
					if !c.NextArg() {
						return nil, c.ArgErr()
					}

					rule.Path = c.Val()

					if c.NextArg() {
						return nil, c.ArgErr()
					}
				case "url":
					if !c.NextArg() {
						return nil, c.ArgErr()
					}

					rule.URL = c.Val()

					if c.NextArg() {
						return nil, c.ArgErr()
					}
				//case "auth-service-name":
				//	if !c.NextArg() {
				//		return nil, c.ArgErr()
				//	}
				//
				//	rule.AuthServiceName = c.Val()
				//
				//	if c.NextArg() {
				//		return nil, c.ArgErr()
				//	}
				//case "service-registry-addr":
				//	if !c.NextArg() {
				//		return nil, c.ArgErr()
				//	}
				//
				//	rule.ServiceRegistryAddr = c.Val()
				//
				//	if c.NextArg() {
				//		return nil, c.ArgErr()
				//	}
				case "except":

					if !c.NextArg() {
						return nil, c.ArgErr()
					}

					method := c.Val()

					if !sliceutils.HasString(internal.HttpMethods, method) {
						return nil, c.ArgErr()
					}

					for c.NextArg() {
						path := c.Val()
						rule.ExclusionRules = append(rule.ExclusionRules, internal.ExclusionRule{Method: method, Path: path})
					}
				}
			}
		default:
			return nil, c.ArgErr()
		}
	}

	if c.Next() {
		return nil, c.ArgErr()
	}

	rule.Handler = NewHandle()

	return rule, nil

}

func NewHandle() http.Handler {
	var gwmux = runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs()
		}),
	)
	var opts = rpc.ClientOptions
	var err error

	for _, r := range []register{
		{
			pb.RegisterTestHandlerFromEndpoint,
			fmt.Sprintf("%s:%d", constants.TestServiceHost, constants.TestServicePort),
		},
	} {
		err = r.f(context.Background(), gwmux, r.endpoint, opts)
		if err != nil {
			err = errors.WithStack(err)
			logger.Error(nil, "Dial [%s] failed %+v", r.endpoint, err)
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return formWrapper(mux)
}

// Ref: https://github.com/grpc-ecosystem/grpc-gateway/issues/7#issuecomment-358569373
func formWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonMap := make(map[string]interface{}, len(r.Form))
			for k, v := range r.Form {
				if len(v) > 0 {
					jsonMap[k] = v[0]
				}
			}
			jsonBody, err := json.Marshal(jsonMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
			r.ContentLength = int64(len(jsonBody))
			r.Header.Set("Content-Type", "application/json")
		}

		h.ServeHTTP(w, r)
	})
}
