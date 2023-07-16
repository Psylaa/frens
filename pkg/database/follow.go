package database

import (
	"github.com/jinzhu/gorm"
)

type FollowRepo struct{ Conn *gorm.DB }

type Follow struct {
	BaseModel
	UserID     string `gorm:"not null"`
	FollowedID string `gorm:"not null"`
}
