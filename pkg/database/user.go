package database

import (
	"github.com/jinzhu/gorm"
)

type UserRepo struct{ Conn *gorm.DB }

type User struct {
	BaseModel
	Username string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
