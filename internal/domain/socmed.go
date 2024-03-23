package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name           string    `json:"name" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" validate:"required"`
	UserID         string    `json:"user_id"`
	User           User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}

func ValidateSocialMedia(s *SocialMedia) error {
	validate := validator.New()

	return validate.Struct(s)
}
