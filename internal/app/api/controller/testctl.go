package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"hal9000/internal/app/api/pkg/ginplus"
	"hal9000/proto/greeter"
)

type TestController struct {
	greeterClient greeter.GreeterService
}


func NewTestController(greeterclient greeter.GreeterService)*TestController{
	return &TestController{greeterClient:greeterclient}
}


func (t *TestController)Hello(c *gin.Context){
	item := c.Param("id")
	if item == ""{
		item = "greeter"
	}
	result, err := t.greeterClient.SayHello(context.TODO(), &greeter.SayRequest{
		Msg: item,
	})
	if err != nil {
		fmt.Println("err is ", err.Error())
		ginplus.ResJSON(c, 200, ginplus.HTTPError{Error: ginplus.HTTPErrorItem{
			Code:    500,
			Message: "rpcerror",
		}})
		return
	}
	fmt.Println(result.Rsp)
	ginplus.ResOK(c)
}
