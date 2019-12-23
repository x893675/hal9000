package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"hal9000/internal/app/api"
	"hal9000/internal/app/api/config"
	"hal9000/internal/app/api/controller"
	"hal9000/internal/app/api/rpcclients/greeter"
	"hal9000/internal/app/api/version"
	k8s "hal9000/pkg/micro/web"
	"log"
	"time"
)



func main(){
	ctx, cancel := context.WithCancel(context.Background())
	// shutdown after 5 seconds
	go func() {
		<-time.After(time.Second * 5)
		log.Println("Shutdown example: shutting down service")
		cancel()
	}()

	service := k8s.NewService(
		web.Name(config.ServiceName),
		web.Version(version.Version),
		web.Context(ctx),
	)
	service.Handle("/", InitWeb())
	_ = service.Init(
		web.Address("0.0.0.0:"+config.ServicePort),
	)

	log.Println("创建服务:名称:" + config.ServiceName + ",版本:" + version.Version)

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