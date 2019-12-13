package micro

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	grpcclient "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/kubernetes"

	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services
	"github.com/micro/go-plugins/client/selector/static"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Selector(s),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpc.NewService(options...)
}


func NewClient(opts ...client.Option) client.Client{
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []client.Option{
		client.Registry(r),
		client.Selector(s),
	}

	options = append(options, opts...)

	return grpcclient.NewClient(options...)
}