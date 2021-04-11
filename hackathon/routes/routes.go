package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hackathon/controller"
	"hackathon/middleware"
)

func Init() {
	r := gin.Default()
	r.Use(cors.Default())
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("api/auth/register", controller.Register)
	r.POST("api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
