package database

import (
	"errors"
	"log"

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

func VerifyUser(username string, password string) (*User, error) {
	// Find the user with the given username
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		log.Println("User tried to login with username:", username, "but it was not found")
		return nil, errors.New("username not found")
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
