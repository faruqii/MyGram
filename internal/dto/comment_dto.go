package dto

import (
	"MyGram/internal/helper"

	"github.com/go-playground/validator/v10"
)

type CommentRequest struct {
	Message string `json:"message" validate:"required" error:"Message is required"`
	PhotoID string `json:"photo_id" validate:"required" error:"Photo ID is required"`
}

// validate Request
func (c *CommentRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[CommentRequest](*c, validate)
}

type CommentResponse struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	PhotoID   string `json:"photo_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type AllCommentResponse struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	PhotoID   string `json:"photo_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
	User      struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	Photo struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   string `json:"user_id"`
	}
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required" error:"Message is required"`
}

// validate Request
func (u *UpdateCommentRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[UpdateCommentRequest](*u, validate)
}

type UpdateCommentResponse struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	PhotoID  string `json:"photo_id"`
	UserID   string `json:"user_id"`
	UpdateAt string `json:"updated_at"`
}
