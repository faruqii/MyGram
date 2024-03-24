package repositories

import (
	"MyGram/internal/domain"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *domain.Comment) error
	FindByID(id string) (*domain.Comment, error)
	GetAll() ([]domain.Comment, error)
	Update(comment *domain.Comment) error
	Delete(id string) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *domain.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) FindByID(id string) (*domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) GetAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) Update(comment *domain.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Comment{}).Error
}
