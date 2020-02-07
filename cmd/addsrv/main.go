package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hal9000/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

var (
	zapLogger  *zap.Logger
	customFunc grpc_zap.CodeToLevel
)


type AddService struct {

}


func (a *AddService) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumReply, error) {
	grpc_ctxtags.Extract(ctx).Set("custom_tags.string", "something").Set("custom_tags.int", 1337)

	// Extract a single request-scoped zap.Logger and log messages. (containing the grpc.xxx tags)
	l := ctxzap.Extract(ctx)
	l.Info("some ping")
	l.Info("another ping")
	return &pb.SumReply{
		V: req.A + req.B,
	},nil
}
func (a *AddService) Concat(ctx context.Context, req *pb.ConcatRequest) (*pb.ConcatReply, error) {
	return &pb.ConcatReply{
		V: req.A + req.B,
	},nil
}

func main()  {

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		//zapOpts := 	[]grpc_zap.Option{
		//	grpc_zap.WithLevels(customFunc),
		//}
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapLogger,_ = cfg.Build(zap.AddCallerSkip(2))
		opts := []grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				//grpc_zap.UnaryServerInterceptor(zapLogger, zapOpts...),
				grpc_zap.UnaryServerInterceptor(zapLogger)),
		}
		server := grpc.NewServer(opts...)
		pb.RegisterAddServer(server, &AddService{})
		listener, err := net.Listen("tcp", ":8891")
		if err != nil {
			log.Fatal(err)
		}
		errc <- server.Serve(listener)
	}()
	log.Println(<-errc)
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}