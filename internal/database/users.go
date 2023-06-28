package database

import (
	"errors"
	"log"

	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username          string         `gorm:"unique" json:"username"`
	Email             string         `json:"email"`
	Bio               string         `json:"bio"`
	Password          string         `json:"-"`
	ProfilePictureURL string         `json:"profilePictureURL"`
	CoverImageURL     string         `json:"coverImageURL"`
	Privacy           shared.Privacy `json:"privacy"`
}

type UserRepo struct {
	db *gorm.DB
}

func (ur *UserRepo) GetUser(id uuid.UUID) (*User, error) {
	var user User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) GetUsers() ([]User, error) {
	var users []User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepo) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) CreateUser(username string, email string, password string) (*User, error) {
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

	if err := ur.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (ur *UserRepo) VerifyUser(username string, password string) (*User, error) {
	user, err := ur.GetUserByUsername(username)
	if err != nil {
		log.Println("User tried to login with username:", username, "but it was not found")
		return nil, errors.New("username not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (ur *UserRepo) UpdateUser(id uuid.UUID, bio string, profilePictureURL string, coverImageURL string) (*User, error) {
	user, err := ur.GetUser(id)
	if err != nil {
		return nil, err
	}

	user.Bio = bio
	user.ProfilePictureURL = profilePictureURL
	user.CoverImageURL = coverImageURL

	if err := ur.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
