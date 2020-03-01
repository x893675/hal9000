package config

import (
	"github.com/emicklei/go-restful"
	"hal9000/pkg/server/runtime"
	"net/http"
	"reflect"
	"strings"
)

func InstallAPI(c *restful.Container) {
	ws := runtime.NewWebService(runtime.GroupVersion{
		Group:   "",
		Version: "v1",
	})

	ws.Route(ws.GET("/configz").
		To(func(request *restful.Request, response *restful.Response) {
			var conf = *sharedConfig

			conf.stripEmptyOptions()

			response.WriteAsJson(convertToMap(&conf))
		}).
		Doc("Get system components configuration").
		Produces(restful.MIME_JSON).
		Writes(Config{}).
		Returns(http.StatusOK, "ok", Config{}))

	c.Add(ws)
}

// convertToMap simply converts config to map[string]bool
// to hide sensitive information
func convertToMap(conf *Config) map[string]bool {
	result := make(map[string]bool, 0)

	if conf == nil {
		return result
	}

	c := reflect.Indirect(reflect.ValueOf(conf))

	for i := 0; i < c.NumField(); i++ {
		name := strings.Split(c.Type().Field(i).Tag.Get("json"), ",")[0]
		if strings.HasPrefix(name, "-") {
			continue
		}

		if c.Field(i).IsNil() {
			result[name] = false
		} else {
			result[name] = true
		}
	}

	return result
}
