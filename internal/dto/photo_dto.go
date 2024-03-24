package dto

import (
	"MyGram/internal/helper"

	"github.com/go-playground/validator/v10"
)

type PhotoRequest struct {
	Title    string `json:"title" validate:"required" error:"Title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required" error:"Photo URL is required"`
}

// validate Request
func (p *PhotoRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[PhotoRequest](*p, validate)
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"required" error:"Title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required" error:"Photo URL is required"`
}

// validate Request
func (u *UpdatePhotoRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[UpdatePhotoRequest](*u, validate)
}

type PhotoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type UpdatePhotoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    string `json:"user_id"`
	UpdatedAt string `json:"created_at"`
}

type AllPhotoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	User      struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
}
