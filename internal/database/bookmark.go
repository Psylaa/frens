package database

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmarks interface {
	Base[Bookmark]
	Get(count int, cursor *time.Time) ([]*Bookmark, error)
	GetByID(id *uuid.UUID) (*Bookmark, error)
	GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]*Bookmark, error)
	GetByPostIDAndUserID(postID *uuid.UUID, userID *uuid.UUID) (*Bookmark, error)
	Exists(userID *uuid.UUID, postID *uuid.UUID) bool
}

type Bookmark struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`
	PostID uuid.UUID `gorm:"type:uuid;not null"`
	Post   Post      `gorm:"foreignKey:PostID"`
}

type BookmarkRepo struct {
	*BaseRepo[Bookmark]
}

func NewBookmarkRepo(db *gorm.DB) Bookmarks {
	return &BookmarkRepo{NewBaseRepo[Bookmark](db)}
}

func (br *BookmarkRepo) Get(count int, cursor *time.Time) ([]*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "Get")

	var bookmarks []*Bookmark
	var result *gorm.DB
	if cursor == nil {
		result = br.db.Limit(count).Find(&bookmarks)
	} else {
		result = br.db.Where("created_at < ?", cursor).Limit(count).Find(&bookmarks)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return bookmarks, nil
}

func (br *BookmarkRepo) GetByID(id *uuid.UUID) (*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "GetByID")

	var bookmark Bookmark
	result := br.db.Preload("User").Preload("Post").First(&bookmark, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmark, nil
}

func (br *BookmarkRepo) GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "GetPaginated")

	var bookmarks []*Bookmark
	result := br.db.Where("user_id = ?", userID).Limit(count).Offset(offset).Find(&bookmarks)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookmarks, nil
}

func (br *BookmarkRepo) GetByPostIDAndUserID(postID *uuid.UUID, userID *uuid.UUID) (*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "GetByPostIDAndUserID")

	var bookmark Bookmark
	result := br.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&bookmark)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmark, nil
}

func (br *BookmarkRepo) Exists(bookmarkID *uuid.UUID, userID *uuid.UUID) bool {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "IsOwner")

	var bookmark Bookmark
	result := br.db.Where("id = ? AND user_id = ?", bookmarkID, userID).First(&bookmark)
	return result.Error == nil
}
