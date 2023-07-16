package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	gorm.Model
	UserID uuid.UUID `gorm:"not null" jsonapi:"attr,userID"`
	PostID uuid.UUID `gorm:"not null" jsonapi:"attr,postID"`
}
