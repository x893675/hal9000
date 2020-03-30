package account

import (
	"context"
)

func (a *AccountService) Builder(ctx context.Context, req interface{}) interface{} {
	//sender := ctxutils.GetSender(ctx)
	//switch r := req.(type) {
	//case *pb.CreatePasswordResetRequest:
	//	r.UserId = pbutil.ToProtoString(sender.UserId)
	//	return r
	//}
	return req
}

func (a *AccountService) Checker(ctx context.Context, req interface{}) error {
	//switch r:= req.(type) {
	//case *pb.DescribeUsersRequest:
	//	return rpc.NewChecker(ctx, r).
	//		Required().
	//		Exec()
	//case *pb.CreateUserRequest:
	//	return rpc.NewChecker(ctx, r).
	//		Required().
	//		Exec()
	//case *pb.GetUserRequest:
	//	return rpc.NewChecker(ctx,r).
	//		Required("user_id").
	//		Exec()
	//}
	return nil
}
