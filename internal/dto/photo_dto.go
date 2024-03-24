package dto

import "github.com/go-playground/validator/v10"

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
}

// validate Request
func (p *PhotoRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type PhotoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
