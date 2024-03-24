package services

import (
	"MyGram/internal/domain"
	"MyGram/internal/repositories"
)

type CommentServices interface {
	Create(comment *domain.Comment) error
	FindByID(id string) (*domain.Comment, error)
	GetAll() ([]domain.Comment, error)
	Update(comment *domain.Comment) error
	Delete(id string) error
}

type commentServices struct {
	commentRepo repositories.CommentRepository
}

func NewCommentServices(commentRepo repositories.CommentRepository) *commentServices {
	return &commentServices{commentRepo: commentRepo}
}

func (s *commentServices) Create(comment *domain.Comment) error {
	err := s.commentRepo.Create(comment)
	if err != nil {
		return HandleError(err, "Failed to create comment", 500)
	}

	return nil
}

func (s *commentServices) FindByID(id string) (*domain.Comment, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return nil, HandleError(err, "Comment not found", 404)
	}

	return comment, nil
}

func (s *commentServices) GetAll() ([]domain.Comment, error) {
	comments, err := s.commentRepo.GetAll()
	if err != nil {
		return nil, HandleError(err, "Failed to get comments", 500)
	}

	return comments, nil
}

func (s *commentServices) Update(comment *domain.Comment) error {
	err := s.commentRepo.Update(comment)
	if err != nil {
		return HandleError(err, "Failed to update comment", 500)
	}

	return nil
}

func (s *commentServices) Delete(id string) error {
	err := s.commentRepo.Delete(id)
	if err != nil {
		return HandleError(err, "Failed to delete comment", 500)
	}

	return nil
}
