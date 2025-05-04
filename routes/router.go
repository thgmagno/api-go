package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thgmagno/api-go/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/shorten-url", controllers.ShortenUrl)
	r.GET("/short", controllers.RedirectToOriginalUrl)
	r.GET("/recently-shortened", controllers.RecentlyShortenedUrls)
	r.GET("/flush-all", controllers.FlushAll)

	return r
}
