package main

import (
	"hal9000/cmd/grpc-gateway/app"
	"log"
)

func main()  {
	cmd := app.NewGrpcGatewayCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}