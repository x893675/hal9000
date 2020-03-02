package main

import (
	"log"
	"net/http"
)

func main() {
	handle := http.StripPrefix("/", http.FileServer(http.Dir("./")))
	log.Println("please open http://localhost:9090/swagger-ui/ to navigate the API in the browser.")
	err := http.ListenAndServe(":9090", handle)
	if err != nil {
		log.Fatal(err)
	}
}
