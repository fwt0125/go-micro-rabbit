package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRoute := gin.Default()
	ginRoute.Use(middleware.Cors(),middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("fwt-secret"))
	ginRoute.Use(sessions.Sessions("mySession", store))
	v1 := ginRoute.Group("/api/v1"){
		v1.GET("ping", func(context *gin.Context) {
			context.JSONP(200, "success")
		})

		v1.POST("user/register", handlers.UserRegiste)
		v1.POST("user/login", handlers.UserLogin)
	}

}