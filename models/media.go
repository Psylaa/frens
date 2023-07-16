package models

import (
	"github.com/jinzhu/gorm"
)

type Media struct {
	gorm.Model
	UserID string `gorm:"not null" jsonapi:"attr,userID"`
	PostID string `gorm:"not null" jsonapi:"attr,postID"`
}
