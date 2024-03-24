package routes

import (
	"MyGram/internal/handlers"
	"MyGram/internal/middleware"
	"MyGram/internal/services"

	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, userSvc services.UserServices, photoSvc services.PhotoServices, middleware *middleware.Middleware) {
	userController := handlers.NewUserHandlers(userSvc, *middleware)
	photoController := handlers.NewPhotoHandlers(photoSvc, *middleware)

	user := router.Group("/users")
	user.Post("/register", userController.Register)
	user.Post("/login", userController.Login)

	// Authenticated routes for users
	userAuth := router.Group("/users")
	userAuth.Use(middleware.Authenticate())
	userAuth.Put("", userController.Update)
	userAuth.Delete("", userController.Delete)

	// Photo routes
	photo := router.Group("/photos")
	photo.Use(middleware.Authenticate())
	photo.Post("", photoController.Create)
	photo.Get("", photoController.GetAll)
	photo.Put("/:id", photoController.Update)
	photo.Delete("/:id", photoController.Delete)
}
