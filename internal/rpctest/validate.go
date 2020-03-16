package rpctest

import (
	"context"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
)

func (p *Server) Checker(ctx context.Context, req interface{}) error {
	switch r:= req.(type) {
	case *pb.DescribeUsersRequest:
		return rpc.NewChecker(ctx, r).
			Required("email").
			Exec()
	case *pb.CreateUserRequest:
		return rpc.NewChecker(ctx, r).
			Required().
			Exec()
	}
	return nil
}

func (p *Server) Builder(ctx context.Context, req interface{}) interface{} {
	//sender := ctxutils.GetSender(ctx)
	//switch r := req.(type) {
	//case *pb.CreatePasswordResetRequest:
	//	r.UserId = pbutil.ToProtoString(sender.UserId)
	//	return r
	//}
	return req
}