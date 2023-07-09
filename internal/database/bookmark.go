package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmarks interface {
	Base[Bookmark]
	GetByID(id *uuid.UUID) (*Bookmark, error)
	IsOwner(bookmarkID *uuid.UUID, userID *uuid.UUID) bool
	GetByPostID(postID *uuid.UUID) (*Bookmark, error)
	GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]*Bookmark, error)
}

type Bookmark struct {
	BaseModel
	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
	PostID uuid.UUID
	Post   Post `gorm:"foreignKey:PostID"`
}

type BookmarkRepo struct {
	*BaseRepo[Bookmark]
}

func NewBookmarkRepo(db *gorm.DB) Bookmarks {
	return &BookmarkRepo{NewBaseRepo[Bookmark](db)}
}

// GetByID returns the bookmark with the given ID
func (br *BookmarkRepo) GetByID(id *uuid.UUID) (*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "GetByID")

	var bookmark Bookmark
	result := br.db.Preload("User").Preload("Post").First(&bookmark, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmark, nil
}

func (br *BookmarkRepo) IsOwner(bookmarkID *uuid.UUID, userID *uuid.UUID) bool {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "IsOwner")

	var bookmark Bookmark
	result := br.db.Where("id = ? AND user_id = ?", bookmarkID, userID).First(&bookmark)
	return result.Error == nil
}

func (br *BookmarkRepo) GetByPostID(postID *uuid.UUID) (*Bookmark, error) {
	logger.DebugLogRequestReceived("database", "BookmarkRepo", "GetByPostID")

	var bookmark Bookmark
	result := br.db.Where("post_id = ?", postID).First(&bookmark)
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
