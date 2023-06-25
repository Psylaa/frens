package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func VerifyUser(username string, password string) (*User, error) {
	// Find the user with the given username
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
