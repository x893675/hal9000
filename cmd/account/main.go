package main

import (
	"hal9000/cmd/account/app"
	"log"
)

func main() {
	cmd := app.NewAccountServiceCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
