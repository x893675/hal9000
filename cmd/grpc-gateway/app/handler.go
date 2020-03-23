package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hal9000/pkg/constants"
	"hal9000/pkg/logger"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const (
	RequestIdKey = "X-Request-Id"
)

type register struct {
	f        func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	endpoint string
}


type Server struct {

}

func log() gin.HandlerFunc {
	l := logger.NewLogger()
	l.HideCallstack()
	return func(c *gin.Context) {
		requestID := ""
		if c.Request.Header.Get(RequestIdKey) != "" {
			requestID = uuid.New().String()
			c.Request.Header.Set(RequestIdKey, requestID)
			c.Writer.Header().Set(RequestIdKey, requestID)
		}
		t := time.Now()

		// process request
		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		logStr := fmt.Sprintf("%s | %3d | %v | %s | %s %s %s",
			requestID,
			statusCode,
			latency,
			clientIP, method,
			path,
			c.Errors.String(),
		)

		switch {
		case statusCode >= 400 && statusCode <= 499:
			l.Warn(nil, logStr)
		case statusCode >= 500:
			l.Error(nil, logStr)
		default:
			l.Info(nil, logStr)
		}
	}
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				logger.Critical(nil, "Panic recovered: %+v\n%s", err, string(httprequest))
				c.JSON(500, gin.H{
					"title": "Error",
					"err":   err,
				})
			}
		}()
		c.Next() // execute all the handlers
	}
}

func (s *Server)run(addr string, port int) error{
	gin.SetMode(gin.ReleaseMode)

	mainHandler := gin.WrapH(s.mainHandler())

	r := gin.New()
	r.Use(log())
	r.Use(recovery())
	r.Any("/v1/*filepath", mainHandler)

	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}

func (s *Server) mainHandler() http.Handler {
	var gwmux = runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD{
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
	}{
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