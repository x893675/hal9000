package rpctest

import (
	"google.golang.org/grpc"
	"hal9000/pkg/constants"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
)

type Server struct {

}

func Serve() {
	s := Server{}
	rpc.NewGrpcServer("test-service", constants.TestServicePort).
		ShowErrorCause(true).
		WithChecker(s.Checker).
		WithBuilder(s.Builder).
		Serve(func(server *grpc.Server) {
			pb.RegisterTestServer(server, &s)
	})
}
