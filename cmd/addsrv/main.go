package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"hal9000/internal/addsrv"
	"hal9000/pb"
	"hal9000/pkg/addr"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	//etcd服务地址
	etcdServer = "etcd:2379"
	ctx   = context.Background()
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

	//etcd服务注册 START
	var host, port string
	var err error
	if cnt := strings.Count(*gRPCAddr, ":"); cnt >= 1 {
		// ipv6 address in format [host]:port or ipv4 host:port
		host, port, err = net.SplitHostPort(*gRPCAddr)
		if err != nil {
			panic(err)
		}
	} else {
		host = *gRPCAddr
	}
	ipAddr, err := addr.Extract(host)
	if err != nil {
		panic(err)
	}
	gaddr := addr.HostPort(ipAddr, port)
	logger.Log("gaddr", gaddr)
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, etcdv3.ClientOptions{})
	if err != nil {
		panic(err)
	}
	prefix := "/hal9000/addsrv/"
	instanceId := uuid.New().String()
	register := etcdv3.NewRegistrar(client,etcdv3.Service{
		Key: prefix + "node-" + instanceId,
		Value: gaddr,
	},logger)
	register.Register()
	defer register.Deregister()
	//etcd服务注册 END

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