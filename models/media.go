package models

import "time"

type Media struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,user"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    uint32    `gorm:"not null" jsonapi:"attr,userID"`
	PostID    uint32    `gorm:"not null" jsonapi:"attr,postID"`
}
