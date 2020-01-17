package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"hal9000/internal/app/api"
	"hal9000/internal/app/api/config"
	"hal9000/internal/app/api/controller"
	"hal9000/internal/app/api/middleware"
	"hal9000/internal/app/api/rpcclients/greeter"
	"hal9000/internal/app/api/trace"
	"hal9000/internal/app/api/version"
	"hal9000/pkg/entrypoint"
	k8s "hal9000/pkg/micro/web"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ep, _ := entrypoint.Initialize()

	_ = trace.New()

	go func() {
		reloadSignal := make(chan os.Signal)
		signal.Notify(reloadSignal, syscall.SIGHUP)
		for {
			sig := <-reloadSignal
			ep.Reload()
			fmt.Printf("OS signaled `%v`, reload", sig.String())
		}
	}()

	go func() {
		gracefulDelay := 50 * time.Millisecond
		shutdownSignal := make(chan os.Signal)
		signal.Notify(shutdownSignal, syscall.SIGTERM, syscall.SIGINT)
		sig := <-shutdownSignal
		fmt.Printf("OS signaled `%v`, graceful shutdown in %s\n", sig.String(), gracefulDelay)
		ctx, _ := context.WithTimeout(context.Background(), gracefulDelay)
		ep.Shutdown(ctx, 0)
	}()

	service := k8s.NewService(
		web.Name(config.ServiceName),
		web.Version(version.Version),
		web.Context(context.Background()),
	)
	service.Handle("/", InitWeb())
	_ = service.Init(
		web.Address("0.0.0.0:" + config.ServicePort),
	)

	log.Println("创建服务:名称:" + config.ServiceName + ",版本:" + version.Version)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitWeb() *gin.Engine {
	gin.SetMode("debug")
	app := gin.Default()
	app.Use(middleware.TraceMiddleware())

	api := CreateApiApplication()

	api.RegisterRouter(app)

	return app
}

func CreateApiApplication() *api.ApiApplication {
	greeterclient := greeter.NewGreeterClient()

	testctl := controller.NewTestController(greeterclient)

	return api.NewApiApplication(testctl)
}
