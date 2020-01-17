package greeter

import (
	"hal9000/internal/app/api/config"
	"hal9000/pkg/micro/micro"
	"hal9000/proto/greeter"
)

//import (
//	"github.com/micro/go-micro/client"
//	"github.com/micro/go-plugins/client/selector/static"
//	"github.com/micro/go-plugins/registry/kubernetes"
//	"github.com/micro/go-plugins/transport/grpc"
//	"hal9000/proto/greeter"
//)
//
//func NewGreeterClient() greeter.GreeterService{
//	k := kubernetes.NewRegistry()
//	st := static.NewSelector()
//	return greeter.NewGreeterService("greeter", client.NewClient(
//		client.Registry(k),
//		client.Selector(st),
//		client.Transport(grpc.NewTransport(),
//	)))
//}

//import (
//	"github.com/micro/go-micro/client"
//	"github.com/micro/go-micro/client/grpc"
//	"github.com/micro/go-plugins/client/selector/static"
//	"github.com/micro/go-plugins/registry/kubernetes"
//	"hal9000/proto/greeter"
//)
//
//func NewGreeterClient()greeter.GreeterService{
//		k := kubernetes.NewRegistry()
//		st := static.NewSelector()
//		return greeter.NewGreeterService("greeter",
//			grpc.NewClient(client.Registry(k),
//			client.Selector(st)),
//		)
//}

func NewGreeterClient() greeter.GreeterService {
	return greeter.NewGreeterService(config.GreeterServiceName, micro.NewClient())
}
