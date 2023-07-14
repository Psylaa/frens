package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Bio      string
	Verified bool `gorm:"default:false"`
}

func (u *User) ToResponse() *Response {
	return &Response{
		Data: []Data{
			{
				Type: Users,
				ID:   u.ID,
			},
		},
	}
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password)) == nil
}
