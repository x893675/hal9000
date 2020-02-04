package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"hal9000/internal/addsrv"
	"hal9000/pb"
	"io"
	"time"
)

func New(conn *grpc.ClientConn) addsrv.AddService {
	sumEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Sum",
		addsrv.EncodeGRPCSumRequest,
		addsrv.DecodeGRPCSumResponse,
		pb.SumReply{},
	).Endpoint()

	concatEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Concat",
		addsrv.EncodeGRPCConcatRequest,
		addsrv.DecodeGRPCConcatResponse,
		pb.ConcatReply{},
	).Endpoint()

	return addsrv.Endpoints{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}
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