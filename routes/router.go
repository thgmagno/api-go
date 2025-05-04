package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thgmagno/api-go/controllers"
	"github.com/thgmagno/api-go/middleware"
	"github.com/thgmagno/api-go/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.RateLimiter(services.Redis, 5, time.Hour))

	r.POST("/shorten-url", controllers.ShortenUrl)
	r.GET("/short", controllers.RedirectToOriginalUrl)
	r.GET("/recently-shortened", controllers.RecentlyShortenedUrls)
	r.GET("/flush-all", controllers.FlushAll)

	return r
}
