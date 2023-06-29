package database

import (
	"errors"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username         string         `gorm:"unique" json:"username"`
	Email            string         `json:"email"`
	Bio              string         `json:"bio"`
	Password         string         `json:"-"`
	ProfilePicture   File           `gorm:"foreignKey:ProfilePictureID" json:"profilePicture"`
	ProfilePictureID uuid.UUID      `json:"-"`
	CoverImage       File           `gorm:"foreignKey:CoverImageID" json:"coverImage"`
	CoverImageID     uuid.UUID      `json:"-"`
	Privacy          shared.Privacy `json:"privacy"`
}

type UserRepo struct {
	db *gorm.DB
}

func (ur *UserRepo) GetUser(id uuid.UUID) (*User, error) {
	var user User
	if err := ur.db.Preload("ProfilePicture").Preload("CoverImage").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) GetUsers() ([]User, error) {
	var users []User
	if err := ur.db.Preload("ProfilePicture").Preload("CoverImage").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := ur.db.Preload("ProfilePicture").Preload("CoverImage").Where("username = ?", username).First(&user).Error; err != nil {
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
		logger.Log.Debug().Str("package", "database").Msgf("Error getting user: %s", err.Error())
		return nil, errors.New("username not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error comparing passwords: %s", err.Error())
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (ur *UserRepo) UpdateUser(id uuid.UUID, bio *string, profilePicture *File, coverImage *File) (*User, error) {
	user, err := ur.GetUser(id)
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error getting user: %s", err.Error())
		return nil, err
	}

	// Create a map to hold the updated fields
	updates := make(map[string]interface{})
	if bio != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Updating bio to %s", *bio)
		updates["bio"] = *bio
	} else {
		logger.Log.Debug().Str("package", "database").Msgf("Bio is nil")
	}
	if profilePicture != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Updating profile picture to %s", profilePicture.ID)
		updates["profile_picture_id"] = profilePicture.ID
	} else {
		logger.Log.Debug().Str("package", "database").Msgf("Profile picture is nil")
	}
	if coverImage != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Updating cover image to %s", coverImage.ID)
		updates["cover_image_id"] = coverImage.ID
	} else {
		logger.Log.Debug().Str("package", "database").Msgf("Cover image is nil")
	}

	// Use the Updates method to only update non-zero values
	if err := ur.db.Model(user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return user, nil
}
