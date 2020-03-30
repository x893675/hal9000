package account

import (
	"context"
	"hal9000/pkg/logger"
	"hal9000/pkg/pb"
	"hal9000/pkg/utils/pbutil"
)

func (a *AccountService) DescribeUsers(context.Context, *pb.DescribeUsersRequest) (*pb.DescribeUsersResponse, error) {
	logger.Info(nil, "in DescribeUsers")
	logger.Info(nil, "")
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
