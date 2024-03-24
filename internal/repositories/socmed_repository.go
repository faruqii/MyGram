package repositories

import (
	"MyGram/internal/domain"

	"gorm.io/gorm"
)

type SocmedRepository interface {
	Create(socmed *domain.SocialMedia) error
	GetAll() ([]domain.SocialMedia, error)
	FindByID(id string) (*domain.SocialMedia, error)
	Update(socmed *domain.SocialMedia) error
	Delete(id string) error
}

type socmedRepository struct {
	db *gorm.DB
}

func NewSocmedRepository(db *gorm.DB) *socmedRepository {
	return &socmedRepository{db: db}
}

func (r *socmedRepository) Create(socmed *domain.SocialMedia) error {
	return r.db.Create(socmed).Error
}

func (r *socmedRepository) GetAll() ([]domain.SocialMedia, error) {
	var socmeds []domain.SocialMedia
	err := r.db.Preload("User").Find(&socmeds).Error
	if err != nil {
		return nil, err
	}

	return socmeds, nil
}

func (r *socmedRepository) FindByID(id string) (*domain.SocialMedia, error) {
	var socmed domain.SocialMedia
	err := r.db.Where("id = ?", id).First(&socmed).Error
	if err != nil {
		return nil, err
	}

	return &socmed, nil
}

func (r *socmedRepository) Update(socmed *domain.SocialMedia) error {
	return r.db.Save(socmed).Error
}

func (r *socmedRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.SocialMedia{}).Error
}