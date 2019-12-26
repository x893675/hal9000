package ginplus

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/metadata"
	ot "github.com/opentracing/opentracing-go"
	"hal9000/internal/app/api/pkg/errors"
	"hal9000/pkg/tracing"
	"hal9000/pkg/util"
	"net/http"
)

// 定义上下文中的键
const (
	prefix = "hal9000"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res_body"
	RootTraceCtx = prefix + "/rootTrace_ctx"
	RequestIDKey = prefix + "/request_id"
)

var jaegerTraceHeaderDict = map[string]string{
	"x-request-id" : "X-Request-Id",
	"x-b3-traceid" : "X-B3-Traceid",
	"x-b3-spanid"  : "X-B3-Spanid",
	"x-b3-sampled" : "X-B3-Sampled",
	"x-b3-parentspanid" : "X-B3-Parentspanid",
	"x-b3-flags"   : "X-B3-Flags",
	"x-ot-span-context" : "X-Ot-Span-Context",
}

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error" swaggo:"true,错误项"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code" swaggo:"true,错误码"`
	Message string `json:"message" swaggo:"true,错误信息"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status" swaggo:"true,状态(OK)"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

func ResOK(c *gin.Context) {
	ResSuccess(c, HTTPStatus{Status: "OK"})
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := util.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}


// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	statusCode := 500
	errItem := HTTPErrorItem{
		Code:    500,
		Message: "服务器发生错误",
	}

	if errCode, ok := errors.FromErrorCode(err.Error()); ok {
		errItem.Code = errCode.Code
		errItem.Message = errCode.Message
		statusCode = errCode.HTTPStatusCode
	}

	if len(status) > 0 {
		statusCode = status[0]
	}

	if statusCode == 500 && err != nil {
		//span := logger.StartSpan(NewContext(c))
		//span = span.WithField("stack", fmt.Sprintf("%+v", err))
		//span.Errorf(err.Error())
		fmt.Println("Err is ", err.Error())
	}

	ResJSON(c, statusCode, HTTPError{Error: errItem})
}

func StartChildSpan(c *gin.Context, operationName string, tags tracing.Tags)  context.Context{
	if rootTranceCtx, ok := c.Get(RootTraceCtx); ok{
		if ctx, ok := rootTranceCtx.(ot.SpanContext); ok{
			span := ot.StartSpan(operationName, ot.ChildOf(ctx), ot.Tags(tags))
			ctx := ot.ContextWithSpan(context.Background(), span)
			return ctx
		}
	}
	return context.Background()
}

func FinishSpan(ctx context.Context) {
	span := ot.SpanFromContext(ctx)
	if span != nil {
		span.Finish()
	}
}

func InjectJaegerTraceToRpcMetaData(c *gin.Context) context.Context{
	md := make(map[string]string)
	for k, v := range jaegerTraceHeaderDict{
		if temp := c.GetHeader(v); temp != "" {
			md[k] = temp
		}
	}
	ctx := metadata.NewContext(context.Background(), md)
	return ctx
}