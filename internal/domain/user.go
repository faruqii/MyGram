package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Username  string    `json:"username" validate:"required,unique"`
	Email     string    `json:"email" validate:"required,email,unique"`
	Password  string    `json:"password" validate:"required,min=6"`
	Age       int       `json:"age" validate:"required,age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return nil
}
