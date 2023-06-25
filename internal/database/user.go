package database

import (
	"github.com/google/uuid"
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

	newUser := User{
		BaseModel: BaseModel{ID: uuid.New()},
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
	}

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}
