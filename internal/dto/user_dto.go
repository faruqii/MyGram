package dto

import "github.com/golang-jwt/jwt/v4"

type RegisterRequest struct {
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

type RegisterResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
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
	ID          string `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
	UpdatedAt   string `json:"updated_at"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
