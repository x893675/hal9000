package greeter

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/client/selector/static"
	"github.com/micro/go-plugins/registry/kubernetes"
	"hal9000/proto/greeter"
)

func NewGreeterClient() greeter.GreeterService{
	k := kubernetes.NewRegistry()
	st := static.NewSelector()
	return greeter.NewGreeterService("greeter", client.NewClient(client.Registry(k),client.Selector(st)))
}