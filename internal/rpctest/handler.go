package rpctest

import (
	"context"
	"hal9000/pkg/logger"
	"hal9000/pkg/pb"
	"hal9000/pkg/utils/pbutil"
	"time"
)

var (
	_    pb.TestServer = (*Server)(nil)

)


func (p *Server) DescribeUsers(ctx context.Context, req *pb.DescribeUsersRequest) (*pb.DescribeUsersResponse, error) {

	logger.Info(nil, "in DescribeUsers")
	logger.Info(nil, "")
	reply := &pb.DescribeUsersResponse{
		TotalCount:           1,
		UserSet:              []*pb.User{
			{
				UserId:               nil,
				Username:             nil,
				Email:                nil,
				PhoneNumber:          nil,
				Description:          pbutil.ToProtoString("description"),
				Status:               pbutil.ToProtoString("1"),
				CreateTime:           nil,
				UpdateTime:           nil,
				StatusTime:           pbutil.ToProtoTimestamp(time.Now()),
			},
		},
	}
	return reply, nil
}


func (p *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	logger.Info(nil, "in CreateUser")

	reply := &pb.CreateUserResponse{
		UserId:               pbutil.ToProtoString("1111"),
	}

	return reply, nil
}