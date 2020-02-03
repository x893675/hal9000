package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"hal9000/internal/addsrv"
	"hal9000/pb"
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
