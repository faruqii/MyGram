package handlers

import (
	"MyGram/internal/domain"
	"MyGram/internal/dto"
	"MyGram/internal/middleware"
	"MyGram/internal/services"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PhotoHandlers struct {
	photoSvc          services.PhotoServices
	middlewareManager middleware.Middleware
}

func NewPhotoHandlers(photoSvc services.PhotoServices, middlewareManager middleware.Middleware) *PhotoHandlers {
	return &PhotoHandlers{
		photoSvc:          photoSvc,
		middlewareManager: middlewareManager,
	}
}

func (c *PhotoHandlers) Create(ctx *fiber.Ctx) error {
	req := dto.PhotoRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	userId := ctx.Locals("user").(string)

	photo := domain.Photo{
		Title:     req.Title,
		Caption:   req.Caption,
		PhotoURL:  req.PhotoURL,
		UserID:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := c.photoSvc.Create(&photo)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := dto.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.Status(http.StatusCreated).JSON(response)

}

// Get All Photos
func (c *PhotoHandlers) GetAll(ctx *fiber.Ctx) error {
	photos, err := c.photoSvc.GetAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := []dto.AllPhotoResponse{}

	for _, photo := range photos {
		res := dto.AllPhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		res.User.Email = photo.User.Email
		res.User.Username = photo.User.Username

		response = append(response, res)
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// Update a Photo
func (c *PhotoHandlers) Update(ctx *fiber.Ctx) error {
	req := dto.UpdatePhotoRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	photoID := ctx.Params("id")
	if photoID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Photo ID is required"})
	}

	// Fetch user ID from token
	userId := ctx.Locals("user").(string)

	// Retrieve photo by ID
	photo, err := c.photoSvc.FindByID(photoID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Photo not found"})
	}

	// Check if the user is authorized to update this photo
	if photo.UserID != userId {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Update photo details
	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoURL = req.PhotoURL
	photo.UpdatedAt = time.Now()

	// Save updated photo
	err = c.photoSvc.Update(photo)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update photo"})
	}

	// Respond with updated photo details
	response := dto.UpdatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// delete a photo
func (c *PhotoHandlers) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Photo ID is required"})
	}

	// Fetch user ID from token
	userId := ctx.Locals("user").(string)

	// Retrieve photo by ID
	photo, err := c.photoSvc.FindByID(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Photo not found"})
	}

	// Check if the user is authorized to delete this photo
	if photo.UserID != userId {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Delete photo
	err = c.photoSvc.Delete(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete photo"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Photo deleted successfully"})
}
