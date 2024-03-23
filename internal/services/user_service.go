package services

import (
	"MyGram/internal/domain"
	"MyGram/internal/dto"
	"MyGram/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
	CreateUserToken(user *domain.User) (string, error)
	Update(user *domain.User) error
	FindUserByID(id string) (*domain.User, error)
	Delete(id string) error
	DeleteUserToken(userID string) error
}

type userServices struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewUserServices(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository) *userServices {
	return &userServices{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *userServices) Register(user *domain.User) error {
	_, err := s.userRepo.FindByEmail(user.Email)
	if err == nil {
		return HandleError(err, "Email already exists", 400)
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return HandleError(err, "Failed to hash password", 500)
	}

	user.Password = string(pwd)

	err = s.userRepo.Create(user)
	if err != nil {
		return HandleError(err, "Failed to create user", 500)
	}

	return nil
}

func (s *userServices) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", HandleError(err, "User not found", 404)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", HandleError(err, "Invalid password", 400)
	}

	token, err := s.CreateUserToken(user)
	if err != nil {
		return "", HandleError(err, "Failed to create token", 500)
	}

	return token, nil
}

func (s *userServices) CreateUserToken(user *domain.User) (string, error) {
	// Create JWT token
	claims := dto.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	// Create or update token in repository
	newToken := &domain.Token{
		UserID: user.ID,
		Token:  signedToken,
	}
	_, err = s.tokenRepo.CreateOrUpdateToken(newToken)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *userServices) Update(user *domain.User) error {
	err := s.userRepo.Update(user)
	if err != nil {
		return HandleError(err, "Failed to update user", 500)
	}

	return nil
}

func (s *userServices) FindUserByID(id string) (*domain.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, HandleError(err, "User not found", 404)
	}

	return user, nil
}

func (s *userServices) Delete(id string) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return HandleError(err, "Failed to delete user", 500)
	}

	return nil
}

func (s *userServices) DeleteUserToken(userID string) error {
	err := s.tokenRepo.DeleteTokenByUserID(userID)
	if err != nil {
		return HandleError(err, "Failed to delete token", 500)
	}

	return nil
}

