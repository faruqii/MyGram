package repositories

import (
	"MyGram/internal/domain"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo *domain.Photo) error
	GetAll() ([]domain.Photo, error)
	FindByID(id string) (*domain.Photo, error)
	Update(photo *domain.Photo) error
	Delete(id string) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) Create(photo *domain.Photo) error {
	return r.db.Create(photo).Error
}

func (r *photoRepository) GetAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	// Preload the user to get the user details
	err := r.db.Preload("User").Find(&photos).Error
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (r *photoRepository) FindByID(id string) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Where("id = ?", id).First(&photo).Error
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

func (r *photoRepository) Update(photo *domain.Photo) error {
	return r.db.Save(photo).Error
}

func (r *photoRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Photo{}).Error
}
