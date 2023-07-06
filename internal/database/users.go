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
	PhoneNumberExists(phone string) bool
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

func NewUser(username string, email string, phoneNumber string, password string) User {
	logger.DebugLogRequestReceived("database", "users", "NewUser")

	// Check if the phoneNumber parameter is empty
	// This is necessary for weirdness with gorm and the unique constraint
	var phonePtr *string
	if phoneNumber != "" {
		phonePtr = &phoneNumber
	}

	// Hash the password
	hashedPass, _ := shared.HashPassword(password)

	return User{
		BaseModel: BaseModel{
			ID: uuid.New(),
		},
		Username:    username,
		Email:       email,
		PhoneNumber: phonePtr,
		Password:    *hashedPass,
		Privacy:     shared.PrivacyPublic,
		Role:        shared.RoleUser,
	}
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

func (ur *UserRepo) PhoneNumberExists(phoneNumber string) bool {
	logger.DebugLogRequestReceived("database", "users", "PhoneNumberExists")
	var user User
	result := ur.db.Where("phone_number = ?", phoneNumber).First(&user)
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
