package repositories

import (
	"MyGram/internal/domain"

	"gorm.io/gorm"
)

type TokenRepository interface {
	CreateOrUpdateToken(token *domain.Token) (string, error)
	FindUserByToken(token string) (*domain.User, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) CreateOrUpdateToken(token *domain.Token) (string, error) {
	var existingToken domain.Token
	err := r.db.Where("user_id = ?", token.UserID).First(&existingToken).Error

	if err != nil {
		if err := r.db.Create(token).Error; err != nil {
			return "", err
		}
	} else {
		token.ID = existingToken.ID
		if err := r.db.Save(token).Error; err != nil {
			return "", err
		}
	}

	return token.Token, nil

}

func (r *tokenRepository) FindUserByToken(token string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
