package api

import "github.com/gin-gonic/gin"

func (api *ApiApplication) RegisterRouter(app *gin.Engine) {
	g := app.Group("/api")


	v1 := g.Group("/v1")
	{
		v1.GET("/hello/:id", api.testCtl.Hello)
	}
}