package database

import (
	"github.com/jinzhu/gorm"
)

type LikeRepo struct{ Conn *gorm.DB }

type Like struct {
	BaseModel
	UserID string `gorm:"not null"`
	PostID string `gorm:"not null"`
}
