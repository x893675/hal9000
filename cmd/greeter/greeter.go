package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	srv "hal9000/internal/app/greeter"
	"hal9000/internal/app/greeter/config"
	"hal9000/internal/app/greeter/version"
	k8s "hal9000/pkg/micro/micro"
	"hal9000/proto/greeter"
	"log"
	"time"
)

func main() {
	service := k8s.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		//micro.Transport(grpc.NewTransport()),
		micro.Version(version.Version),
	)

	_ = service.Server().Init(
		server.Wait(nil),
	)
	service.Init(
		micro.AfterStart(func() error {
			log.Println("after start...")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("after stop...")
			return nil
		}),
	)

	greeterSrv := srv.NewGreeterServer()

	err := greeter.RegisterGreeterHandler(service.Server(), greeterSrv)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("创建服务:名称:" + config.ServiceName + ",版本:" + version.Version)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
