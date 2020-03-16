package main

import (
	"hal9000/cmd/rpctest/app"
	"log"
)

func main() {
	cmd := app.NewTestServerCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}