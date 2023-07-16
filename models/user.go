package models

import "time"

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" jsonapi:"primary,user"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	Username  string    `gorm:"not null;unique" jsonapi:"attr,username"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
}
