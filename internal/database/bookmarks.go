package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BookmarkRepository struct{ db *gorm.DB }

func (br *BookmarkRepository) Create(bookmark *models.Bookmark) error {
	return br.db.Create(bookmark).Error
}

func (br *BookmarkRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	query := br.db

	if limit != nil {
		query = query.Limit(*limit)
	}

	if cursor != nil {
		query = query.Where("created_at < ?", cursor)
	}

	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	err := query.Find(&bookmarks).Error
	return bookmarks, err
}

func (br *BookmarkRepository) Update(bookmark *models.Bookmark) error {
	return br.db.Save(bookmark).Error
}

func (br *BookmarkRepository) Delete(id uuid.UUID) error {
	return br.db.Where("id = ?", id).Delete(&models.Bookmark{}).Error
}

func (br *BookmarkRepository) GetBySourceID(id uuid.UUID) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	err := br.db.Where("source_id = ?", id).Find(&bookmarks).Error
	return bookmarks, err
}

func (br *BookmarkRepository) GetByTargetID(id uuid.UUID) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	err := br.db.Where("target_id = ?", id).Find(&bookmarks).Error
	return bookmarks, err
}
