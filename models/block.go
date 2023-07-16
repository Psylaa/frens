package models

import (
	"time"

	"github.com/google/uuid"
)

type Block struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,user"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    uuid.UUID `gorm:"not null" jsonapi:"attr,userID"`
	BlockedID uuid.UUID `gorm:"not null" jsonapi:"attr,blockedID"`
}
