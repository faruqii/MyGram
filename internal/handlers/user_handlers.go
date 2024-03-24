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

type UserHandlers struct {
	userSvc           services.UserServices
	middlewareManager middleware.Middleware
}

func NewUserHandlers(userSvc services.UserServices, middlewareManager middleware.Middleware) *UserHandlers {
	return &UserHandlers{
		userSvc:           userSvc,
		middlewareManager: middlewareManager,
	}
}

func (c *UserHandlers) Register(ctx *fiber.Ctx) error {
	req := dto.RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := domain.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		Age:       req.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := c.userSvc.Register(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

func (c *UserHandlers) Login(ctx *fiber.Ctx) error {
	req := dto.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := c.userSvc.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.LoginResponse{
		Token: token,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (c *UserHandlers) Update(ctx *fiber.Ctx) error {
	// Get user ID from context
	token := ctx.Locals("user").(string)

	// Get user from repository
	user, err := c.userSvc.FindUserByID(token)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := dto.UpdateUserRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user.Email = req.Email
	user.Username = req.Username
	user.UpdatedAt = time.Now()

	err = c.userSvc.Update(user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.UpdateUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (c *UserHandlers) Delete(ctx *fiber.Ctx) error {
	// Get user ID from context
	token := ctx.Locals("user").(string)

	// Delete user token
	err := c.userSvc.DeleteUserToken(token)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.userSvc.Delete(token)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Your account has been deleted successfully"})
}
