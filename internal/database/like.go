package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type LikeRepository struct{ db *gorm.DB }

func (r *LikeRepository) Create(like *models.Like) (*models.Like, error) {
	if err := r.db.Create(like).Error; err != nil {
		return nil, err
	}

	// Reload the like so that we can preload the user and post.
	var loadedLike models.Like
	if err := r.db.
		Preload("User").
		Preload("Post").
		First(&loadedLike, "id = ?", like.ID).
		Error; err != nil {
		return nil, err
	}

	return &loadedLike, nil
}

func (r *LikeRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.Like, error) {
	var likes []models.Like
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

	err := query.Find(&likes).Error
	return likes, err
}

func (r *LikeRepository) GetBySourceID(id uuid.UUID) ([]models.Like, error) {
	var likes []models.Like
	err := r.db.Where("source_id = ?", id).Find(&likes).Error
	return likes, err
}

func (r *LikeRepository) GetByTargetID(id uuid.UUID) ([]models.Like, error) {
	var likes []models.Like
	err := r.db.Where("target_id = ?", id).Find(&likes).Error
	return likes, err
}
