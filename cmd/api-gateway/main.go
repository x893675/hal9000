package main

import (
	"hal9000/cmd/api-gateway/app"
	"log"
)

func main() {
	cmd := app.NewAPIGatewayCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}