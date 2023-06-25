package database

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username string `gorm:"unique" json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func CreateUser(username string, email string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
