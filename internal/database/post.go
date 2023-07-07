package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Posts represents the interface for a Post repository
type Posts interface {
	Base[Post]
}

// Post struct represents the post table in the database with appropriate gorm tags.
type Post struct {
	BaseModel
	AuthorID uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"` // UUID of the post author
	Author   User           `gorm:"foreignKey:AuthorID" json:"author"`   // The author of the post
	Privacy  shared.Privacy `gorm:"type:varchar(20)" json:"privacy"`     // Privacy setting of the post
	Text     string         `gorm:"type:text" json:"text"`               // Text content of the post
	Media    []*File        `gorm:"foreignKey:PostID" json:"media"`      // Associated media files
	MediaIDs []string       `gorm:"-" json:"-"`                          // Helper field to hold the Media ID's while processing a request
}

// PostRepo struct represents the Post repository
type PostRepo struct {
	*BaseRepo[Post]
}

// NewPostRepo initializes and returns a Post repository
func NewPostRepo(db *gorm.DB) Posts {
	return &PostRepo{NewBaseRepo[Post](db)}
}

// GetByID returns the post with the given ID, preloading the Author and Media data
func (pr *PostRepo) GetByID(id *uuid.UUID) (*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByID")

	var post Post
	result := pr.db.Preload("Author").Preload("Media").First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}
