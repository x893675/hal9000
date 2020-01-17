package middleware

import (
	"github.com/gin-gonic/gin"
	ot "github.com/opentracing/opentracing-go"
	"hal9000/internal/app/api/pkg/ginplus"
	"hal9000/pkg/util"
	"log"
)

func TraceMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		if ot.GlobalTracer() == nil {
			requestID := c.GetHeader("X-Request-Id")
			if requestID == "" {
				requestID = util.MustUUID()
			}
			c.Set(ginplus.RequestIDKey, requestID)
			c.Next()
			return
		}
		inBoundSpanCtx, err := ot.GlobalTracer().Extract(ot.HTTPHeaders, ot.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			log.Println(err.Error())
			c.Next()
			return
		}
		c.Set(ginplus.RequestIDKey, c.GetHeader("X-Request-Id"))
		c.Set(ginplus.RootTraceCtx, inBoundSpanCtx)
	}
}
