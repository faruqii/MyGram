package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"

	"MyGram/internal/helper"
)

type RegisterRequest struct {
	Age      int    `json:"age" validate:"required,numeric,min=9" error:"Age must be a numeric value and must be at least 9"`
	Email    string `json:"email" validate:"required,email" error:"Email is required and must be a valid email address"`
	Password string `json:"password" validate:"required,min=6" error:"Password is required and must be at least 6 characters long"`
	Username string `json:"username" validate:"required" error:"Username is required"`
}

func (r *RegisterRequest) Validate(validate *validator.Validate) error {
	return helper.ValidateFunc[RegisterRequest](*r, validate)
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"date_of_birth"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Age       int    `json:"date_of_birth"`
	UpdatedAt string `json:"updated_at"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
