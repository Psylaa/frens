package models

import "time"

type Bookmark struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,bookmark"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    uint32    `gorm:"not null"`
	User      *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" jsonapi:"relation,user"`
	PostID    uint32    `gorm:"not null"`
	Post      *Post     `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" jsonapi:"relation,post"`
}
