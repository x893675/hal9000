package auth

import (
	"context"
	"hal9000/pkg/pb"
	"hal9000/pkg/rpc"
)

func (a *AuthhService) Builder(ctx context.Context, req interface{}) interface{} {
	//sender := ctxutils.GetSender(ctx)
	//switch r := req.(type) {
	//case *pb.CreatePasswordResetRequest:
	//	r.UserId = pbutil.ToProtoString(sender.UserId)
	//	return r
	//}
	return req
}

func (a *AuthhService) Checker(ctx context.Context, req interface{}) error {
	switch r := req.(type) {
	case *pb.AuthRequest:
		return rpc.NewChecker(ctx, r).
			Required("username", "password").
			Exec()
	case *pb.Oauth2Request:
		return rpc.NewChecker(ctx, r).
			Required().
			Exec()
		//case *pb.GetUserRequest:
		//	return rpc.NewChecker(ctx,r).
		//		Required("user_id").
		//		Exec()
	}
	return nil
}
