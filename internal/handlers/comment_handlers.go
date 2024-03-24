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

type CommentHandlers struct {
	commentService    services.CommentServices
	middlewareManager middleware.Middleware
}

func NewCommentHandlers(commentService services.CommentServices, middlewareManager middleware.Middleware) *CommentHandlers {
	return &CommentHandlers{
		commentService:    commentService,
		middlewareManager: middlewareManager,
	}
}

func (c *CommentHandlers) Create(ctx *fiber.Ctx) error {
	req := dto.CommentRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	userID := ctx.Locals("user").(string)

	comment := domain.Comment{
		UserID:    userID,
		PhotoID:   req.PhotoID,
		Message:   req.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := c.commentService.Create(&comment)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := dto.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)

}

func (c *CommentHandlers) GetAll(ctx *fiber.Ctx) error {
	comments, err := c.commentService.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := []dto.AllCommentResponse{}

	for _, comment := range comments {
		commentResponse := dto.AllCommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:  comment.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		commentResponse.User.ID = comment.User.ID
		commentResponse.User.Email = comment.User.Email
		commentResponse.User.Username = comment.User.Username

		commentResponse.Photo.ID = comment.Photo.ID
		commentResponse.Photo.Title = comment.Photo.Title
		commentResponse.Photo.Caption = comment.Photo.Caption
		commentResponse.Photo.PhotoURL = comment.Photo.PhotoURL
		commentResponse.Photo.UserID = comment.Photo.UserID

		response = append(response, commentResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *CommentHandlers) Update(ctx *fiber.Ctx) error {
	req := dto.UpdateCommentRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	validate := validator.New()

	if err := req.Validate(validate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	commentID := ctx.Params("id")

	if commentID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Comment ID is required"})
	}

	comment, err := c.commentService.FindByID(commentID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	userID := ctx.Locals("user").(string)

	if comment.UserID != userID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	comment.Message = req.Message
	comment.UpdatedAt = time.Now()

	err = c.commentService.Update(comment)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := dto.UpdateCommentResponse{
		ID:       comment.ID,
		Message:  comment.Message,
		PhotoID:  comment.PhotoID,
		UserID:   comment.UserID,
		UpdateAt: comment.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *CommentHandlers) Delete(ctx *fiber.Ctx) error {
	commentID := ctx.Params("id")

	if commentID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Comment ID is required"})
	}

	comment, err := c.commentService.FindByID(commentID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Comment not found"})
	}

	userID := ctx.Locals("user").(string)

	if comment.UserID != userID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = c.commentService.Delete(commentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Comment deleted successfully"})
}
