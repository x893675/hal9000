package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/transport/grpc"
	srv "hal9000/internal/app/greeter"
	k8s "hal9000/pkg/micro/micro"
	"hal9000/proto/greeter"
	"log"
	"time"
)



func main(){
	service := k8s.NewService(
		micro.Name("greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Transport(grpc.NewTransport()),
	)
	service.Init()

	server := srv.NewGreeterServer()

	err := greeter.RegisterGreeterHandler(service.Server(), server)
	if err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}