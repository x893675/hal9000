package auth

import (
	"google.golang.org/grpc"
	"hal9000/pkg/constants"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
)

type AuthhService struct {
}

func Serve() {
	s := AuthhService{}
	rpc.NewGrpcServer("auth-service", constants.AuthServicePort).
		ShowErrorCause(true).
		WithChecker(s.Checker).
		WithBuilder(s.Builder).
		Serve(func(server *grpc.Server) {
			pb.RegisterAuthServer(server, &s)
		})
}
