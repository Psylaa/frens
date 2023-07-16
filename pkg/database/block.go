package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BlockRepo struct{ Conn *gorm.DB }

type Block struct {
	BaseModel
	UserID    uuid.UUID `gorm:"not null"`
	BlockedID uuid.UUID `gorm:"not null"`
}
