package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"hal9000/internal/app/api"
	"hal9000/internal/app/api/controller"
	"hal9000/internal/app/api/rpcclients/greeter"
	k8s "hal9000/pkg/micro/web"
)



func main(){
	service := k8s.NewService(
		web.Name("hal9000-api"),
	)
	service.Handle("/", InitWeb())
	_ = service.Init(
		web.Address("0.0.0.0:8080"),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitWeb() *gin.Engine {
	gin.SetMode("debug")

	//app := gin.New()
	app := gin.Default()
	//app.NoMethod(middleware.NoMethodHandler())
	//app.NoRoute(middleware.NoRouteHandler())
	//// 崩溃恢复
	//app.Use(middleware.RecoveryMiddleware())
	api := CreateApiApplication()

	api.RegisterRouter(app)

	return app
}

func CreateApiApplication()*api.ApiApplication{
	greeterclient := greeter.NewGreeterClient()

	testctl := controller.NewTestController(greeterclient)

	return api.NewApiApplication(testctl)
}