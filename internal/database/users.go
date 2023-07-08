package database

import (
	"errors"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Users interface {
	Base[User]
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	UsernameExists(username string) bool
	EmailExists(email string) bool
	IsVerifiedByID(id *uuid.UUID) (bool, error)
	CheckCredentials(username, password string) (*User, error)
}

type User struct {
	BaseModel
	Username    string  `gorm:"unique"`
	Email       string  `gorm:"unique"`
	PhoneNumber *string `gorm:"unique,default:null"` // Need default null for unique constraint to work
	Password    string

	Bio *string `gorm:"size:1024"`

	Privacy shared.Privacy
	Role    shared.Role

	Avatar   File `gorm:"foreignKey:AvatarID"`
	Cover    File `gorm:"foreignKey:CoverID"`
	CoverID  *uuid.UUID
	AvatarID *uuid.UUID

	VerifiedEmail       bool
	VerifiedPhoneNumber bool
}

func (u *User) VerifyPassword(password string) bool {
	// Use bcrypt's CompareHashAndPassword to compare the provided password with the hashed one
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type UserRepo struct {
	*BaseRepo[User]
}

func NewUserRepo(db *gorm.DB) Users {
	return &UserRepo{NewBaseRepo[User](db)}
}

func (ur *UserRepo) GetByEmail(email string) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "GetByEmail")
	var user User
	result := ur.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (ur *UserRepo) GetByUsername(username string) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "GetByUsername")
	var user User
	result := ur.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (ur *UserRepo) UsernameExists(username string) bool {
	logger.DebugLogRequestReceived("database", "users", "UsernameExists")
	var user User
	result := ur.db.Where("username = ?", username).First(&user)
	return !result.RecordNotFound()
}

func (ur *UserRepo) EmailExists(email string) bool {
	logger.DebugLogRequestReceived("database", "users", "EmailExists")
	var user User
	result := ur.db.Where("email = ?", email).First(&user)
	return !result.RecordNotFound()
}

// IsVerifiedByID checks if a user with the given ID has verified their email or phone number.
func (ur *UserRepo) IsVerifiedByID(id *uuid.UUID) (bool, error) {
	logger.DebugLogRequestReceived("database", "users", "IsVerifiedByID")
	user, err := ur.GetByID(id)
	if err != nil {
		return false, err
	}

	return user.VerifiedEmail || user.VerifiedPhoneNumber, nil
}

func (ur *UserRepo) CheckCredentials(username, password string) (*User, error) {
	logger.DebugLogRequestReceived("database", "users", "CheckCredentials")
	user, err := ur.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if !user.VerifyPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (ur *UserRepo) UpdatePassword(id *uuid.UUID, password string) error {
	logger.DebugLogRequestReceived("database", "users", "UpdatePassword")
	hashedPass, err := shared.HashPassword(password)
	if err != nil {
		return err
	}

	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	user.Password = *hashedPass
	return ur.db.Save(user).Error
}

func (ur *UserRepo) UpdateEmail(id *uuid.UUID, email string) error {
	logger.DebugLogRequestReceived("database", "users", "UpdateEmail")
	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	user.Email = email
	return ur.db.Save(user).Error
}

func (ur *UserRepo) UpdatePhoneNumber(id *uuid.UUID, phoneNumber string) error {
	logger.DebugLogRequestReceived("database", "users", "UpdatePhoneNumber")
	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	user.PhoneNumber = &phoneNumber
	return ur.db.Save(user).Error
}

func (ur *UserRepo) UpdateBio(id *uuid.UUID, bio string) error {
	logger.DebugLogRequestReceived("database", "users", "UpdateBio")
	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	user.Bio = &bio
	return ur.db.Save(user).Error
}

func (ur *UserRepo) UpdatePrivacy(id *uuid.UUID, privacy shared.Privacy) error {
	logger.DebugLogRequestReceived("database", "users", "UpdatePrivacy")
	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	user.Privacy = privacy
	return ur.db.Save(user).Error
}
