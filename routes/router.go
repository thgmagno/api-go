package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thgmagno/api-go/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/shorten-url", controllers.ShortenUrl)
	r.GET("/short", controllers.RedirectToOriginalUrl)

	return r
}
