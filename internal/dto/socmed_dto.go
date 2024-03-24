package dto

import (
	"MyGram/internal/helper"

	"github.com/go-playground/validator/v10"
)

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required" error:"Name is required"`
	SocialMediaURL string `json:"social_media_url" validate:"required" error:"Social Media URL is required"`
}

func (s *SocialMediaRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[SocialMediaRequest](*s, validate)
}

type SocialMediaResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         string `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required" error:"Name is required"`
	SocialMediaURL string `json:"social_media_url" validate:"required" error:"Social Media URL is required"`
}

func (u *UpdateSocialMediaRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[UpdateSocialMediaRequest](*u, validate)
}

type UpdateSocialMediaResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         string `json:"user_id"`
	UpdatedAt      string `json:"updated_at"`
}

type AllSocialMediaResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         string `json:"user_id"`
	CreatedAt      string `json:"created_at"`
	UpdateAt       string `json:"updated_at"`
	User           struct {
		ID       string `json:"id"`
		Usernmae string `json:"username"`
		Email    string `json:"email"`
	}
}
