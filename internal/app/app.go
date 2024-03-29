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
	photoRepo := repositories.NewPhotoRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	socmedRepo := repositories.NewSocmedRepository(db)

	userSvc := services.NewUserServices(userRepo, tokenRepo)
	photoSvc := services.NewPhotoServices(photoRepo)
	commentSvc := services.NewCommentServices(commentRepo)
	socmedSvc := services.NewSocmedServices(socmedRepo)

	middleware := middleware.NewMiddleware(tokenRepo)

	api := app.Group("/api")
	routes.Routes(api, userSvc, photoSvc, commentSvc, socmedSvc, middleware)

	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}

}
