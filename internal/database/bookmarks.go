package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	BaseModel
	UserID uuid.UUID
	PostID uuid.UUID
	Owner  User `gorm:"foreignKey:UserID"`
}

type BookmarkRepo struct {
	db *gorm.DB
}

func (br *BookmarkRepo) Create(bookmark *Bookmark) error {
	logger.DebugLogRequestReceived("database", "bookmark", "Create")

	if err := br.db.Create(bookmark).Error; err != nil {
		return err
	}

	return nil
}

func (br *BookmarkRepo) DeleteByID(bookmarkID *uuid.UUID) error {
	logger.DebugLogRequestReceived("database", "bookmark", "Delete")

	if err := br.db.Delete(&Bookmark{}, "id = ?", bookmarkID).Error; err != nil {
		return err
	}

	return nil
}

func (br *BookmarkRepo) ExistsByPostAndUserID(postID *uuid.UUID, userID *uuid.UUID) (bool, error) {
	logger.DebugLogRequestReceived("database", "bookmark", "ExistsByPostAndUserID")

	var bookmark Bookmark
	if err := br.db.
		Where("post_id = ? AND user_id = ?", postID, userID).
		First(&bookmark).
		Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (br *BookmarkRepo) GetByID(bookmarkID *uuid.UUID) (*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "bookmark", "GetByID")

	var bookmark Bookmark
	if err := br.db.
		Preload("Owner").
		Where("id = ?", bookmarkID).
		First(&bookmark).
		Error; err != nil {
		return nil, err
	}

	return &bookmark, nil
}

func (br *BookmarkRepo) GetByPostID(postID *uuid.UUID, count *int, offset *int) ([]*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "bookmark", "GetByPostID")

	var bookmarks []*Bookmark
	query := br.db.
		Preload("Owner").
		Where("post_id = ?", postID)

	// If count is provided, add it to the query
	if count != nil {
		query = query.Limit(*count)
	}

	// If offset is provided, add it to the query
	if offset != nil {
		query = query.Offset(*offset)
	}

	// Execute the query
	if err := query.Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (br *BookmarkRepo) GetByPostAndUserID(userID *uuid.UUID, postID *uuid.UUID) (*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "bookmark", "GetByUserAndPostID")

	var bookmark Bookmark
	if err := br.db.
		Preload("Owner").
		Where("user_id = ? AND post_id = ?", userID, postID).
		First(&bookmark).
		Error; err != nil {
		return nil, err
	}

	return &bookmark, nil
}

func (br *BookmarkRepo) GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "bookmark", "GetByUserID")

	var bookmarks []*Bookmark
	query := br.db.
		Preload("Owner").
		Where("user_id = ?", userID)

	// If count is provided, add it to the query
	if count != nil {
		query = query.Limit(*count)
	}

	// If offset is provided, add it to the query
	if offset != nil {
		query = query.Offset(*offset)
	}

	// Execute the query
	if err := query.Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}
