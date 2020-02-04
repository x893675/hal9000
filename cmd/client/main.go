package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"hal9000/internal/addsrv"
	"time"
)

func main() {

	addService := NewAddService()
	println("Sum Response:")
	sum(context.Background(), addService, 11111, 22222)

	println("Concat Response:")
	concat(context.Background(), addService, "11111", "22222")

	//gRPCAddr := flag.String("gRPC", ":8891", "gRPC client")
	//flag.Parse()
	//
	//ctx, cancel := context.WithTimeout(context.Background(),
	//	10000*time.Millisecond)
	//defer cancel()
	//
	//conn, err := grpc.DialContext(ctx,
	//	*gRPCAddr, grpc.WithInsecure(),
	//)
	//
	//if err != nil {
	//	//log.Fatalln("gRPC dial error:", err)
	//	logger.Log("gRPC dial error", err)
	//}
	//defer conn.Close()
	//
	//addService := New(conn)
	//
	//println("Sum Response:")
	//sum(context.Background(), addService, 11111, 22222)
	//
	//println("Concat Response:")
	//concat(context.Background(), addService, "11111", "22222")
}

func NewAddService() addsrv.AddService {
	//etcd sd start
	etcdServer := "127.0.0.1:2379"
	client, err := etcdv3.NewClient(context.Background(), []string{etcdServer}, etcdv3.ClientOptions{})
	if err != nil {
		panic(err)
	}
	logger :=log.NewNopLogger()
	prefix := "/hal9000/addsrv/"
	instancer, err := etcdv3.NewInstancer(client, prefix, logger)
	if err != nil {
		panic(err)
	}
	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, SumEndpointFactory, logger)
	//创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)

	sumEndPoint := lb.Retry(3, 3*time.Second, balancer)

	endpointer = sd.NewEndpointer(instancer, ConcatEndpointFactory, logger)
	//创建负载均衡器
	balancer = lb.NewRoundRobin(endpointer)

	contactEndPoint := lb.Retry(3, 3*time.Second, balancer)

	return addsrv.Endpoints{
		SumEndpoint:    sumEndPoint,
		ConcatEndpoint: contactEndPoint,
	}
	//etcd sd end
}


func sum(ctx context.Context, svc addsrv.AddService, a, b int) {
	output := svc.Sum(ctx, a, b)
	println(output)
}

func concat(ctx context.Context, svc addsrv.AddService, a, b string) {
	output := svc.Concat(ctx, a, b)
	println(output)
}