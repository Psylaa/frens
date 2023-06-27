package database

import (
	"errors"
	"log"

	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username       string         `gorm:"unique" json:"username"`
	Email          string         `json:"email"`
	Bio            string         `json:"bio"`
	Password       string         `json:"-"`
	ProfilePicture *string        `json:"profilePicture"`
	BannerImage    *string        `json:"bannerImage"`
	Privacy        shared.Privacy `json:"privacy"`
}

func GetUsers() ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id uuid.UUID) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
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
		Privacy:   shared.PrivacyPublic,
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

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id uuid.UUID, bio *string, profilePicture *string, bannerImage *string) (*User, error) {

	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if bio != nil {
		user.Bio = *bio
	}

	if profilePicture != nil {
		user.ProfilePicture = profilePicture
	}

	if bannerImage != nil {
		user.BannerImage = bannerImage
	}

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
