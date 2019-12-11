package greeter

import (
	"github.com/micro/go-micro/client"
	"hal9000/proto/greeter"
)

func NewGreeterClient() greeter.GreeterService{
	return greeter.NewGreeterService("go.micro.srv.greeter", client.DefaultClient)
}