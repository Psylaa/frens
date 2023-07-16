package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Block struct {
	gorm.Model
	UserID    uuid.UUID `gorm:"not null" jsonapi:"attr,userID"`
	BlockedID uuid.UUID `gorm:"not null" jsonapi:"attr,blockedID"`
}
