package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type MediaRepository struct{ db *gorm.DB }

func (m *MediaRepository) Create(media *models.Media) (*models.Media, error) {
	if err := m.db.Create(media).Error; err != nil {
		return nil, err
	}

	// Reload the media so that we can preload the user and post.
	// At the moment, media is added to a post AFTER the upload.
	// However, for posterity reasons, we'll preload the post as well.
	var loadedMedia models.Media
	if err := m.db.
		Preload("User").
		Preload("Post").
		First(&loadedMedia, "id = ?", media.ID).
		Error; err != nil {
		return nil, err
	}

	return &loadedMedia, nil
}

func (m *MediaRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.Media, error) {
	var media []models.Media
	query := m.db

	if limit != nil {
		query = query.Limit(*limit)
	}

	if cursor != nil {
		query = query.Where("created_at < ?", cursor)
	}

	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	err := query.Find(&media).Error
	return media, err
}

func (m *MediaRepository) Update(media *models.Media) error {
	return m.db.Save(media).Error
}

func (m *MediaRepository) Delete(id uuid.UUID) error {
	return m.db.Where("id = ?", id).Delete(&models.Media{}).Error
}
