package database

import (
	"github.com/jinzhu/gorm"
)

type PostRepo struct{ Conn *gorm.DB }

type Post struct {
	BaseModel
	UserID string `gorm:"not null"`
	Text   string `gorm:"not null"`
}
