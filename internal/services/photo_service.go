package services

import (
	"MyGram/internal/domain"
	"MyGram/internal/repositories"
)

type PhotoServices interface {
	Create(photo *domain.Photo) error
	GetAll() ([]domain.Photo, error)
	FindByID(id string) (*domain.Photo, error)
	Update(photo *domain.Photo) error
	Delete(id string) error
}

type photoServices struct {
	photoRepo repositories.PhotoRepository
}

func NewPhotoServices(photoRepo repositories.PhotoRepository) *photoServices {
	return &photoServices{photoRepo: photoRepo}
}

func (s *photoServices) Create(photo *domain.Photo) error {
	err := s.photoRepo.Create(photo)
	if err != nil {
		return HandleError(err, "Failed to create photo", 500)
	}

	return nil
}

func (s *photoServices) GetAll() ([]domain.Photo, error) {
	photos, err := s.photoRepo.GetAll()
	if err != nil {
		return nil, HandleError(err, "Failed to get photos", 500)
	}

	return photos, nil
}

func (s *photoServices) FindByID(id string) (*domain.Photo, error) {
	photo, err := s.photoRepo.FindByID(id)
	if err != nil {
		return nil, HandleError(err, "Photo not found", 404)
	}

	return photo, nil
}

func (s *photoServices) Update(photo *domain.Photo) error {
	err := s.photoRepo.Update(photo)
	if err != nil {
		return HandleError(err, "Failed to update photo", 500)
	}

	return nil
}

func (s *photoServices) Delete(id string) error {
	err := s.photoRepo.Delete(id)
	if err != nil {
		return HandleError(err, "Failed to delete photo", 500)
	}

	return nil
}
