package app

import (
	"MyGram/internal/config"
	"MyGram/internal/middleware"
	"MyGram/internal/repositories"
	"MyGram/internal/routes"
	"MyGram/internal/services"

	"github.com/gofiber/fiber/v2"
)

func StartApplication() {
	app := fiber.New()

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)

	userSvc := services.NewUserServices(userRepo, tokenRepo)
	middleware := middleware.NewMiddleware(tokenRepo)
	api := app.Group("/api")
	routes.Routes(api, userSvc, middleware)
	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}

}
