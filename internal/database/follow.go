package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type FollowRepository struct{ db *gorm.DB }

func (fr *FollowRepository) Create(follow *models.Follow) error {
	return fr.db.Create(follow).Error
}

func (fr *FollowRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.Follow, error) {
	var follows []models.Follow
	query := fr.db

	if limit != nil {
		query = query.Limit(*limit)
	}

	if cursor != nil {
		query = query.Where("created_at < ?", cursor)
	}

	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	err := query.Find(&follows).Error
	return follows, err
}

func (fr *FollowRepository) Update(follow *models.Follow) error {
	return fr.db.Save(follow).Error
}

func (fr *FollowRepository) GetBySourceID(id uuid.UUID) ([]models.Follow, error) {
	var follows []models.Follow
	err := fr.db.Where("source_id = ?", id).Find(&follows).Error
	return follows, err
}

func (fr *FollowRepository) GetByTargetID(id uuid.UUID) ([]models.Follow, error) {
	var follows []models.Follow
	err := fr.db.Where("target_id = ?", id).Find(&follows).Error
	return follows, err
}
