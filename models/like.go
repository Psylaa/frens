package models

import "time"

type Like struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,like"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    uint32    `gorm:"not null" jsonapi:"attr,userID"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PostID    uint32    `gorm:"not null" jsonapi:"attr,postID"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
