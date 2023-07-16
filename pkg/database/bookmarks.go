package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BookmarkRepo struct{ Conn *gorm.DB }

type Bookmark struct {
	BaseModel
	UserID uuid.UUID `gorm:"not null"`
	PostID uuid.UUID `gorm:"not null"`
}
