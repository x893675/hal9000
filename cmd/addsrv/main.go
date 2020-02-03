package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"hal9000/internal/addsrv"
	"hal9000/pb"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	httpAddr := flag.String("HTTP", ":8890", "HTTP server")
	gRPCAddr := flag.String("gRPC", ":8891", "gRPC server")
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.TimestampFormat(
		func() time.Time { return time.Now().Local() },
		time.RFC3339Nano,
	))
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger.Log("msg", "Server Start...")
	defer logger.Log("msg", "Server Stop...")


	svc := addsrv.New()
	endpoints := addsrv.Endpoints{
		SumEndpoint:    addsrv.MakeSumEndpoint(svc),
		ConcatEndpoint: addsrv.MakeConcatEndpoint(svc),
	}

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", *httpAddr)
		handle := addsrv.MakeHTTPHandler(endpoints)
		errc <- http.ListenAndServe(*httpAddr, handle)
	}()

	go func() {
		logger := log.With(logger, "transport", "GRPC")
		logger.Log("addr", *gRPCAddr)
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errc <- err
			return
		}
		srv := addsrv.MakeGRPCServer(endpoints)
		s:= grpc.NewServer()
		pb.RegisterAddServer(s, srv)
		errc <- s.Serve(listener)
	}()

	logger.Log("exit", <-errc)
}