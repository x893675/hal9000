package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hal9000/internal/app/api/pkg/ginplus"
	"hal9000/pkg/tracing"
	"hal9000/proto/greeter"
)

type TestController struct {
	greeterClient greeter.GreeterService
}

func NewTestController(greeterclient greeter.GreeterService) *TestController {
	return &TestController{greeterClient: greeterclient}
}

func (t *TestController) Hello(c *gin.Context) {
	ctx := ginplus.StartChildSpan(c, "TestController.Hello", tracing.Tags{
		"api": "/api/v1/hello/:id",
	})
	item := c.Param("id")
	if item == "" {
		item = "greeter"
	}

	result, err := t.greeterClient.SayHello(ginplus.InjectJaegerTraceToRpcMetaData(c), &greeter.SayRequest{
		Msg: item,
	})
	if err != nil {
		fmt.Println("err is ", err.Error())
		ginplus.FinishSpan(ctx)
		ginplus.ResJSON(c, 200, ginplus.HTTPError{Error: ginplus.HTTPErrorItem{
			Code:    500,
			Message: err.Error(),
		}})
		return
	}
	fmt.Println(result.Rsp)
	ginplus.FinishSpan(ctx)
	ginplus.ResOK(c)
}
