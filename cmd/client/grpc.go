package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"hal9000/internal/addsrv"
	"hal9000/pb"
	"io"
	"time"
)


func NewAddService() addsrv.AddService {
	//etcd sd start
	ctx, cancel := context.WithTimeout(context.Background(),
		3*time.Second)
	defer cancel()
	etcdServer := "127.0.0.1:2379"
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, etcdv3.ClientOptions{})
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

func SumEndpointFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		10000*time.Millisecond)
	defer cancel()
	conn, err := grpc.DialContext(ctx,
		instanceAddr, grpc.WithInsecure(),
	)
	//conn, err := grpc.Dial(instanceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	sumEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Sum",
		addsrv.EncodeGRPCSumRequest,
		addsrv.DecodeGRPCSumResponse,
		pb.SumReply{},
	).Endpoint()
	return sumEndpoint, conn, nil
}

func ConcatEndpointFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		10000*time.Millisecond)
	defer cancel()
	conn, err := grpc.DialContext(ctx,
		instanceAddr, grpc.WithInsecure(),
	)
	//conn, err := grpc.Dial(instanceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	concatEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Concat",
		addsrv.EncodeGRPCConcatRequest,
		addsrv.DecodeGRPCConcatResponse,
		pb.ConcatReply{},
	).Endpoint()
	return concatEndpoint, conn, nil
}