package controllers

import (
	"MyGram/internal/domain"
	"MyGram/internal/dto"
	"MyGram/internal/services"
	"github.com/gin-gonic/gin"
)

type UserControllers struct {
	userSvc services.UserServices
}

func NewUserControllers(userSvc services.UserServices) *UserControllers {
	return &UserControllers{userSvc: userSvc}
}

func (c *UserControllers) Register(ctx *gin.Context) {
	req := dto.RegisterRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		DateOfBirth: req.DateOfBirth,
	}

	err := c.userSvc.Register(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.RegisterResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
	}

	ctx.JSON(201, response)
}

func (c *UserControllers) Login(ctx *gin.Context) {
	req := dto.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userSvc.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.LoginResponse{
		Token: token,
	}

	ctx.JSON(200, response)
}
