package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	_ "hal9000/internal/apiserver/resource/install"
	"hal9000/pkg/constants"
	"hal9000/pkg/server/runtime"
	"io/ioutil"
	"log"
)

var output string

func init() {
	flag.StringVar(&output, "output", "./swagger.json", "--output=./swagger.json")
}

func main() {
	flag.Parse()
	generateSwaggerJson()
}

func generateSwaggerJson() {

	container := runtime.Container

	apiTree(container)

	config := restfulspec.Config{
		WebServices:                   container.RegisteredWebServices(),
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}

	swagger := restfulspec.BuildSwagger(config)

	swagger.Info.Extensions = make(spec.Extensions)
	swagger.Info.Extensions.Add("x-tagGroups", []struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}{
		{
			Name: "Resources",
			Tags: []string{constants.TestResourcesTag},
		},
	})

	data, _ := json.MarshalIndent(swagger, "", "  ")
	err := ioutil.WriteFile(output, data, 420)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successfully written to %s", output)
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "api server",
			Description: "api server OpenAPI",
			Contact: &spec.ContactInfo{
				Name:  "api server",
				Email: "example@example.com",
				URL:   "http://localhost:8080",
			},
			License: &spec.License{
				Name: "Apache",
				URL:  "http://www.apache.org/licenses/",
			},
			Version: "1.0.0",
		},
	}

	// setup security definitions
	swo.SecurityDefinitions = map[string]*spec.SecurityScheme{
		"jwt": spec.APIKeyAuth("Authorization", "header"),
	}
	swo.Security = []map[string][]string{{"jwt": []string{}}}
}

func apiTree(container *restful.Container) {
	buf := bytes.NewBufferString("\n")
	for _, ws := range container.RegisteredWebServices() {
		for _, route := range ws.Routes() {
			buf.WriteString(fmt.Sprintf("%s %s\n", route.Method, route.Path))
		}
	}
	log.Println(buf.String())
}
