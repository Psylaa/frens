package models

import "time"

type Post struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,post"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	UserID    uint32    `gorm:"not null" jsonapi:"attr,userID"`
	Text      string    `gorm:"not null" jsonapi:"attr,text"`
	Privacy   string    `gorm:"not null" jsonapi:"attr,privacy"`
	User      *User     `gorm:"foreignkey:UserID;" jsonapi:"relation,user"`
}
