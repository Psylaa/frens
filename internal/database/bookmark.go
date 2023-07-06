package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmarks interface {
	Base[Bookmark]
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
