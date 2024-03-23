package repositories

import (
	"MyGram/internal/domain"

	"gorm.io/gorm"
)

type TokenRepository interface {
	CreateOrUpdateToken(token *domain.Token) (string, error)
	FindUserByToken(token string) (string, error)
	DeleteTokenByUserID(userID string) error
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

func (r *tokenRepository) FindUserByToken(token string) (string, error) {
	var user domain.Token
	err := r.db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.UserID, nil
}

func (r *tokenRepository) DeleteTokenByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&domain.Token{}).Error
}
