package app

import (
	"MyGram/internal/config"
	"MyGram/internal/repositories"
	"MyGram/internal/routes"
	"MyGram/internal/services"

	"github.com/gin-gonic/gin"
)

func StartApplication() {
	app := gin.Default()

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)

	userSvc := services.NewUserServices(userRepo, tokenRepo)

	api := app.Group("/api")
	{
		routes.Routes(api, userSvc)
	}

	err = app.Run(":3000")

	if err != nil {
		panic(err)
	}

}
