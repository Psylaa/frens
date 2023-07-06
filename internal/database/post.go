package database

import (
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
