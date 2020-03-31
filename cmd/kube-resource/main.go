package main

import (
	"hal9000/cmd/kube-resource/app"
	"log"
)

func main() {

	cmd := app.NewKubeResourceServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}