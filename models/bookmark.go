package models

import (
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,user"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    uuid.UUID `gorm:"not null" jsonapi:"attr,userID"`
	PostID    uuid.UUID `gorm:"not null" jsonapi:"attr,postID"`
}
