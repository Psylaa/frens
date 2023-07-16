package database

import (
	"github.com/jinzhu/gorm"
)

type MediaRepo struct{ Conn *gorm.DB }

type Media struct {
	BaseModel
	UserID string `gorm:"not null"`
	PostID string `gorm:"not null"`
}
