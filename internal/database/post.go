package database

import (
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Posts interface {
	Base[Post]
}

// Post struct represents the post table in the database with appropriate gorm tags.
type Post struct {
	BaseModel
	UserID       uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"` // ID of the user who created the post
	User         User           `json:"user"`                              // User who created the post
	Privacy      shared.Privacy `gorm:"type:varchar(20)" json:"privacy"`   // Privacy setting of the post
	Text         string         `gorm:"type:text" json:"text"`             // Text content of the post
	MediaIDs     []*uuid.UUID   `gorm:"-" json:"-"`                        // Helper field to hold the Media ID's while processing a request
	IsLiked      bool           `gorm:"-" json:"isLiked"`                  // Indicates if post is liked by user
	IsBookmarked bool           `gorm:"-" json:"isBookmarked"`             // Indicates if post is bookmarked by user
}

type PostsRepo struct {
	*BaseRepo[Post]
}

func NewPostRepo(db *gorm.DB) Posts {
	return &PostsRepo{NewBaseRepo[Post](db)}
}
