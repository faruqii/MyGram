package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	PhotoID   string    `json:"photo_id"`
	Photo     Photo     `json:"photo" gorm:"foreignKey:PhotoID;references:ID"`
	Message   string    `json:"message" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}

func ValidateComment(c *Comment) error {
	validate := validator.New()

	return validate.Struct(c)
}
