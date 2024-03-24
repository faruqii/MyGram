package handlers

import (
	"MyGram/internal/domain"
	"MyGram/internal/dto"
	"MyGram/internal/middleware"
	"MyGram/internal/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SocmedHandler struct {
	socmedService     services.SocmedServices
	middlewareManager middleware.Middleware
}

func NewSocmedHandler(socmedService services.SocmedServices, middlewareManager middleware.Middleware) *SocmedHandler {
	return &SocmedHandler{
		socmedService:     socmedService,
		middlewareManager: middlewareManager,
	}
}

func (s *SocmedHandler) Create(ctx *fiber.Ctx) error {
	req := dto.SocialMediaRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	userId := ctx.Locals("user").(string)

	socmed := domain.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         userId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.socmedService.Create(&socmed)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := dto.SocialMediaResponse{
		ID:             socmed.ID,
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         socmed.UserID,
		CreatedAt:      socmed.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": response})
}

func (s *SocmedHandler) GetAll(ctx *fiber.Ctx) error {
	socmed, err := s.socmedService.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := []dto.AllSocialMediaResponse{}

	for _, v := range socmed {
		res := dto.AllSocialMediaResponse{
			ID:             v.ID,
			Name:           v.Name,
			SocialMediaURL: v.SocialMediaURL,
			UserID:         v.UserID,
			CreatedAt:      v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:       v.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		res.User.ID = v.User.ID
		res.User.Usernmae = v.User.Username
		res.User.Email = v.User.Email

		response = append(response, res)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}

func (s *SocmedHandler) Update(ctx *fiber.Ctx) error {
	req := dto.UpdateSocialMediaRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	socmedID := ctx.Params("id")

	socmed, err := s.socmedService.FindByID(socmedID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	userID := ctx.Locals("user").(string)

	if socmed.UserID != userID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	socmed.Name = req.Name
	socmed.SocialMediaURL = req.SocialMediaURL
	socmed.UpdatedAt = time.Now()

	err = s.socmedService.Update(socmed)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := dto.UpdateSocialMediaResponse{
		ID:             socmed.ID,
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         socmed.UserID,
		UpdatedAt:      socmed.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *SocmedHandler) Delete(ctx *fiber.Ctx) error {
	socmedID := ctx.Params("id")

	if socmedID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Social Media ID is required"})
	}

	socmed, err := s.socmedService.FindByID(socmedID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Social Media not found"})
	}

	userID := ctx.Locals("user").(string)

	if socmed.UserID != userID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = s.socmedService.Delete(socmedID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Social Media deleted successfully"})
}
