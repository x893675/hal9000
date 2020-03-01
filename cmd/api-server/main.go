package main

import (
	"hal9000/cmd/api-server/app"
	"log"
)

func main() {

	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}