package services

import (
	"MyGram/internal/domain"
	"MyGram/internal/repositories"
)

type SocmedServices interface {
	Create(socmed *domain.SocialMedia) error
	GetAll() ([]domain.SocialMedia, error)
	FindByID(id string) (*domain.SocialMedia, error)
	Update(socmed *domain.SocialMedia) error
	Delete(id string) error
}

type socmedServices struct {
	socmedRepo repositories.SocmedRepository
}

func NewSocmedServices(socmedRepo repositories.SocmedRepository) *socmedServices {
	return &socmedServices{socmedRepo: socmedRepo}
}

func (s *socmedServices) Create(socmed *domain.SocialMedia) error {
	err := s.socmedRepo.Create(socmed)
	if err != nil {
		return HandleError(err, "Failed to create social media", 500)
	}

	return nil
}

func (s *socmedServices) GetAll() ([]domain.SocialMedia, error) {
	socmeds, err := s.socmedRepo.GetAll()
	if err != nil {
		return nil, HandleError(err, "Failed to get social medias", 500)
	}

	return socmeds, nil
}

func (s *socmedServices) FindByID(id string) (*domain.SocialMedia, error) {
	socmed, err := s.socmedRepo.FindByID(id)
	if err != nil {
		return nil, HandleError(err, "Social Media not found", 404)
	}

	return socmed, nil
}

func (s *socmedServices) Update(socmed *domain.SocialMedia) error {
	err := s.socmedRepo.Update(socmed)
	if err != nil {
		return HandleError(err, "Failed to update social media", 500)
	}

	return nil
}

func (s *socmedServices) Delete(id string) error {
	err := s.socmedRepo.Delete(id)
	if err != nil {
		return HandleError(err, "Failed to delete social media", 500)
	}

	return nil
}
