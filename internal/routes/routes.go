package routes

import (
	"MyGram/internal/controllers"
	"MyGram/internal/services"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, userSvc services.UserServices) {
	userController := controllers.NewUserControllers(userSvc)

	user := router.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}
}
