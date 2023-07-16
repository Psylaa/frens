package models

import (
	"github.com/jinzhu/gorm"
)

type Follow struct {
	gorm.Model
	UserID     string `gorm:"not null" jsonapi:"attr,userID"`
	FollowedID string `gorm:"not null" jsonapi:"attr,followedID"`
}
