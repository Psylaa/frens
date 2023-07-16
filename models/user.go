package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" jsonapi:"attr,username"`
	Email    string `gorm:"not null;unique" jsonapi:"attr,email"`
	Password string `gorm:"not null" jsonapi:"attr,password"`
}
