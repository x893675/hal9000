package account

import (
	"context"
	"hal9000/pkg/logger"
	"hal9000/pkg/pb"
	"hal9000/pkg/utils/ctxutils"
	"hal9000/pkg/utils/pbutil"
)

func (a *AccountService) DescribeUsers(ctx context.Context, req *pb.DescribeUsersRequest) (*pb.DescribeUsersResponse, error) {
	requestId := ctxutils.GetRequestId(ctx)
	logger.Info(ctx, "request id is [%v]", requestId)

	logger.Info(ctx, "in DescribeUsers")
	logger.Info(ctx, "")
	reply := &pb.DescribeUsersResponse{
		TotalCount: 1,
		Users: []*pb.User{
			{
				UserId:      nil,
				Username:    nil,
				Email:       nil,
				PhoneNumber: nil,
				Description: pbutil.ToProtoString("description"),
				Status:      pbutil.ToProtoInt32(0),
				CreateTime:  nil,
				UpdateTime:  nil,
			},
		},
	}
	return reply, nil
}
