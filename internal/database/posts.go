package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type PostRepository struct{ db *gorm.DB }

func (r *PostRepository) Create(post *models.Post) (*models.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}

	// Reload the post so that we can preload the user.
	var loadedPost models.Post
	if err := r.db.
		Preload("User").
		First(&loadedPost, "id = ?", post.ID).
		Error; err != nil {
		return nil, err
	}

	return &loadedPost, nil
}

func (r *PostRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.Post, error) {
	var posts []models.Post
	query := r.db

	if limit != nil {
		query = query.Limit(*limit)
	}

	if cursor != nil {
		query = query.Where("created_at < ?", cursor)
	}

	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	err := query.
		Preload("User").
		Find(&posts).Error
	return posts, err
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&models.Post{}).Error
}
