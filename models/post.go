package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserID string `gorm:"not null" jsonapi:"attr,userID"`
	Text   string `gorm:"not null" jsonapi:"attr,text"`
}
