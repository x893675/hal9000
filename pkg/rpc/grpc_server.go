package rpc

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"hal9000/pkg/grpcerr"
	"hal9000/pkg/logger"
	"hal9000/pkg/utils/ctxutils"
	"hal9000/pkg/version"
	"net"
	"runtime/debug"
	"strings"
	"time"
)

type checkerT func(ctx context.Context, req interface{}) error
type builderT func(ctx context.Context, req interface{}) interface{}

var (
	defaultChecker checkerT
	defaultBuilder builderT
)

type RegisterCallback func(*grpc.Server)

type GrpcServer struct {
	ServiceName    string
	Port           int
	showErrorCause bool
	checker        checkerT
	builder        builderT
}

func NewGrpcServer(serviceName string, port int) *GrpcServer {
	return &GrpcServer{
		ServiceName:    serviceName,
		Port:           port,
		showErrorCause: false,
		checker:        defaultChecker,
		builder:        defaultBuilder,
	}
}

func (g *GrpcServer) ShowErrorCause(b bool) *GrpcServer {
	g.showErrorCause = b
	return g
}

func (g *GrpcServer) WithChecker(c checkerT) *GrpcServer {
	g.checker = c
	return g
}

func (g *GrpcServer) WithBuilder(b builderT) *GrpcServer {
	g.builder = b
	return g
}

func (g *GrpcServer) Serve(callback RegisterCallback, opt ...grpc.ServerOption) {
	logger.Info(nil, "service verson is %s", version.Version)

	logger.Info(nil, "Service [%s] start listen at port [%d]", g.ServiceName, g.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		err = errors.WithStack(err)
		logger.Critical(nil, "failed to listen: %+v", err)
	}

	builtinOptions := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc_middleware.WithUnaryServerChain(
			grpc_validator.UnaryServerInterceptor(),
			g.unaryServerLogInterceptor(),
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				//ctx = db.NewContext(ctx, g.mysqlConfig)

				if g.checker != nil {
					err = g.checker(ctx, req)
					if err != nil {
						return
					}
				}

				return handler(ctx, req)
			},
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				if g.builder != nil {
					req = g.builder(ctx, req)
				}
				return handler(ctx, req)
			},
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
					logger.Critical(nil, "GRPC server recovery with error: %+v", p)
					logger.Critical(nil, string(debug.Stack()))
					if e, ok := p.(error); ok {
						return grpcerr.NewWithDetail(nil, grpcerr.Internal, e, grpcerr.ErrorInternalError)
					}
					return grpcerr.New(nil, grpcerr.Internal, grpcerr.ErrorInternalError)
				}),
			),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
					logger.Critical(nil, "GRPC server recovery with error: %+v", p)
					logger.Critical(nil, string(debug.Stack()))
					if e, ok := p.(error); ok {
						return grpcerr.NewWithDetail(nil, grpcerr.Internal, e, grpcerr.ErrorInternalError)
					}
					return grpcerr.New(nil, grpcerr.Internal, grpcerr.ErrorInternalError)
				}),
			),
		),
	}
	grpcServer := grpc.NewServer(append(opt, builtinOptions...)...)
	reflection.Register(grpcServer)
	callback(grpcServer)

	if err = grpcServer.Serve(lis); err != nil {
		err = errors.WithStack(err)
		logger.Critical(nil, "%+v", err)
	}
}

var (
	jsonPbMarshaller = &jsonpb.Marshaler{
		OrigName: true,
	}
)

func (g *GrpcServer) unaryServerLogInterceptor() grpc.UnaryServerInterceptor {
	showErrorCause := g.showErrorCause

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		requestId := ctxutils.GetRequestId(ctx)
		ctx = ctxutils.SetRequestId(ctx, requestId)
		locale := ctxutils.GetLocale(ctx)
		ctx = ctxutils.SetLocale(ctx, locale)

		method := strings.Split(info.FullMethod, "/")
		action := method[len(method)-1]
		if p, ok := req.(proto.Message); ok {
			if content, err := jsonPbMarshaller.MarshalToString(p); err != nil {
				logger.Error(ctx, "Failed to marshal proto message to string [%s] [%+v]", action, err)
			} else {
				logger.Info(ctx, "Request received [%s] [%s]", action, content)
			}
		}
		start := time.Now()

		resp, err := handler(ctx, req)

		elapsed := time.Since(start)
		logger.Info(ctx, "Handled request [%s] exec_time is [%s]", action, elapsed)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.OK {
				logger.Debug(ctx, "Response is error: %s, %s", e.Code().String(), e.Message())
				if !showErrorCause {
					err = grpcerr.ClearErrorCause(err)
				}
			}
		}
		return resp, err
	}
}