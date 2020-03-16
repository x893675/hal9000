package grpcerr

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hal9000/pkg/logger"
	"hal9000/pkg/pb"
)

//const En = "en"
//const ZhCN = "zh_cn"
//const DefaultLocale = En

func newStatus(ctx context.Context, code codes.Code, err error, errMsg ErrorMessage, a ...interface{}) *status.Status {
	//locale := ctxutil.GetLocale(ctx)
	//if len(locale) == 0 {
	//	locale = DefaultLocale
	//}

	s := status.New(code, errMsg.Message( err, a...))

	errorDetail := &pb.ErrorDetail{ErrorName: errMsg.Name}
	if err != nil {
		errorDetail.Cause = fmt.Sprintf("%+v", err)
	}
	logger.NewLogger().WithDepth(5).Error(ctx, "err: %+v, errMsg: %s", err, errMsg.Message(err, a...))

	sd, e := s.WithDetails(errorDetail)
	if e == nil {
		return sd
	} else {
		logger.NewLogger().WithDepth(5).Error(ctx, "%+v", errors.WithStack(e))
	}
	return s
}

func ClearErrorCause(err error) error {
	if e, ok := status.FromError(err); ok {
		details := e.Details()
		if len(details) > 0 {
			detail := details[0]
			if d, ok := detail.(*pb.ErrorDetail); ok {
				d.Cause = ""
				// clear detail
				proto := e.Proto()
				proto.Details = proto.Details[:0]
				e = status.FromProto(proto)
				e, _ := e.WithDetails(d)
				return e.Err()
			}
		}
	}
	return err
}

type GRPCError interface {
	error
	GRPCStatus() *status.Status
}

func New(ctx context.Context, code codes.Code, errMsg ErrorMessage, a ...interface{}) GRPCError {
	return newStatus(ctx, code, nil, errMsg, a...).Err().(GRPCError)
}

func NewWithDetail(ctx context.Context, code codes.Code, err error, errMsg ErrorMessage, a ...interface{}) GRPCError {
	return newStatus(ctx, code, err, errMsg, a...).Err().(GRPCError)
}

func IsGRPCError(err error) bool {
	if e, ok := err.(GRPCError); ok && e != nil {
		return true
	}
	return false
}