package account

import (
	"google.golang.org/grpc"
	"hal9000/pkg/constants"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
)

type AccountService struct {
}

func Serve() {
	s := AccountService{}
	rpc.NewGrpcServer("account-service", constants.AccountServicePort).
		ShowErrorCause(true).
		WithChecker(s.Checker).
		WithBuilder(s.Builder).
		Serve(func(server *grpc.Server) {
			pb.RegisterAccountServer(server, &s)
		})
}
