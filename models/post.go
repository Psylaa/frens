package models

import "time"

type Post struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,user"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    string    `gorm:"not null" jsonapi:"attr,userID"`
	Text      string    `gorm:"not null" jsonapi:"attr,text"`
}
