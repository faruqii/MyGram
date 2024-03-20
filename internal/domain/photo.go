package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Photo struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" validate:"required"`
	UserID    string    `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

func ValidatePhoto(p *Photo) error {
	validate := validator.New()

	return validate.Struct(p)
}
