package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username    string    `json:"username" validate:"required,unique"`
	Email       string    `json:"email" validate:"required,email,unique"`
	Password    string    `json:"password" validate:"required,min=6"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

func ValidateUser(u *User) error {
	validate := validator.New()

	return validate.Struct(u)
}
