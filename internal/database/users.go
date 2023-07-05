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
	Username    string `gorm:"unique"`
	Email       string
	Bio         string
	Password    string
	Avatar      File `gorm:"foreignKey:AvatarID"`
	AvatarID    uuid.UUID
	Cover       File `gorm:"foreignKey:CoverID"`
	CoverID     uuid.UUID
	Privacy     shared.Privacy
	Role        shared.Role
	IsFollowing bool `gorm:"-"`
}

type UserRepo struct {
	db      *gorm.DB
	Follows *FollowRepo
}

func (ur *UserRepo) GetByID(requestorID, toLookupID *uuid.UUID) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "GetByID")

	var user User
	if err := ur.db.
		Preload("Avatar").
		Preload("Cover").
		Where("id = ?", toLookupID).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	isFollowing, err := ur.Follows.Exists(requestorID, toLookupID)
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error checking if user is following: %s", err.Error())
		return nil, err
	}
	user.IsFollowing = isFollowing
	return &user, nil
}

func (ur *UserRepo) GetUsers() ([]*User, error) {
	logger.DebugLogRequestReceived("database", "users", "GetUsers")

	var users []*User
	if err := ur.db.Preload("Avatar").Preload("Cover").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) GetUserByUsername(username *string) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "GetUserByUsername")

	var user User
	if err := ur.db.Preload("Avatar").Preload("Cover").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) CreateUser(username, email, password string) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "CreateUser")

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

func (ur *UserRepo) VerifyUser(username, password *string) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "VerifyUser")

	user, err := ur.GetUserByUsername(username)
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error getting user: %s", err.Error())
		return nil, errors.New("username not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*password)); err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error comparing passwords: %s", err.Error())
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (ur *UserRepo) UpdateBio(userID *uuid.UUID, bio *string) error {
	logger.DebugLogRequestReceived("database", "users", "UpdateBio")

	user, err := ur.GetByID(userID, userID) // Only the requesting user can update their own bio, so we pass the user's ID as the requestor ID
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error getting user: %s", err.Error())
		return err
	}

	if bio != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Updating bio to %s", *bio)
		if err := ur.db.Model(user).Update("bio", *bio).Error; err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRepo) UpdateAvatar(userID, profilePictureID *uuid.UUID) error {
	logger.DebugLogRequestReceived("database", "users", "UpdateAvatar")

	user, err := ur.GetByID(userID, userID) // Only the requesting user can update their own profile picture, so we pass the user's ID as the requestor ID
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error getting user: %s", err.Error())
		return err
	}

	if profilePictureID != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Updating profile picture to %s", profilePictureID)

		var newAvatar File
		if err := ur.db.First(&newAvatar, "id = ?", profilePictureID).Error; err != nil {
			logger.Log.Debug().Str("package", "database").Msgf("Profile picture not found: %s", err.Error())
			return err
		}

		if err := ur.db.Model(&newAvatar).Updates(File{ID: *profilePictureID}).Error; err != nil {
			return err
		}

		user.Avatar = newAvatar
		if err := ur.db.Save(&user).Error; err != nil {
			return err
		}
	} else {
		logger.Log.Debug().Str("package", "database").Msgf("Profile picture is nil")
		user.AvatarID = uuid.Nil
		if err := ur.db.Save(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRepo) UpdateCover(userID, coverID *uuid.UUID) error {
	logger.DebugLogRequestReceived("database", "users", "UpdateCover")

	user, err := ur.GetByID(userID, userID) // Only the requesting user can update their own cover image, so we pass the user's ID as the requestor ID
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error getting user: %s", err.Error())
		return err
	}
	logger.Log.Debug().Str("package", "database").Msgf("Retrieved user for cover image update: %s", user.ID)

	if coverID != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Updating cover image to %s", coverID)

		var newCover File
		if err := ur.db.First(&newCover, "id = ?", coverID).Error; err != nil {
			logger.Log.Debug().Str("package", "database").Msgf("Cover image not found: %s", err.Error())
			return err
		}

		if err := ur.db.Model(&newCover).Updates(File{ID: *coverID}).Error; err != nil {
			return err
		}

		user.Cover = newCover
		if err := ur.db.Save(&user).Error; err != nil {
			return err
		}
	} else {
		logger.Log.Debug().Str("package", "database").Msgf("Cover image is nil")
		user.CoverID = uuid.Nil
		if err := ur.db.Save(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRepo) UsernameExists(username *string) bool {
	logger.DebugLogRequestReceived("database", "users", "UsernameExists")

	var count int64
	ur.db.Model(&User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (ur *UserRepo) Exists(id *uuid.UUID) bool {
	logger.DebugLogRequestReceived("database", "users", "Exists")

	var count int64
	ur.db.Model(&User{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func (ur *UserRepo) EmailExists(email *string) bool {
	logger.DebugLogRequestReceived("database", "users", "EmailExists")

	var count int64
	ur.db.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (ur *UserRepo) Delete(id *uuid.UUID) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "DeleteUser")

	// Define the user struct
	var user User

	// Retrieve the user before deleting
	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Delete the user
	if err := ur.db.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
