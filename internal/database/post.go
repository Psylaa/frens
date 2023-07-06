package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Posts interface {
	Base[Post]
}

type Post struct {
	BaseModel
	Author   User `gorm:"foreignKey:AuthorID"`
	AuthorID uuid.UUID
	Privacy  shared.Privacy
	Text     string
	Media    []*File `gorm:"foreignKey:PostID;AssociationForeignKey:ID" json:"media"`
}

type PostRepo struct {
	*BaseRepo[Post]
}

func NewPostRepo(db *gorm.DB) Posts {
	return &PostRepo{NewBaseRepo[Post](db)}
}

// GetByID returns the post with the given ID
func (pr *PostRepo) GetByID(id *uuid.UUID) (*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByID")

	var post Post
	result := pr.db.Preload("Author").Preload("Media").First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}
