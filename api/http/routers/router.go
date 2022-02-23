package routers

import (
	"dante-api/api/http/controllers"
	"dante-api/configs"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	gin.SetMode(configs.Config.Env)

	api := router.Group("api")

	userApi := api.Group("user")
	{
		userController := new(controllers.UserController)
		// Store User
		userApi.POST("", userController.Store)
		// Query User
		userApi.GET("", userController.Show)
		// Delete User
		userApi.DELETE("", userController.Delete)
	}

	return router
}
