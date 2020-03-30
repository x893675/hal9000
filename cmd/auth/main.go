package main

import (
	"hal9000/cmd/auth/app"
	"log"
)

func main()  {
	cmd := app.NewAuthServiceCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
