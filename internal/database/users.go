package database

import (
	"bytes"
	"crypto/rand"
	"errors"

	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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
	Salt        []byte

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
	hashedPassword, err := shared.HashPassword(password, u.Salt)
	if err != nil {
		return false
	}
	return bytes.Equal([]byte(*hashedPassword), []byte(u.Password))
}

type UserRepo struct {
	*BaseRepo[User]
}

func NewUserRepo(db *gorm.DB) Users {
	return &UserRepo{NewBaseRepo[User](db)}
}

func NewUser(username string, email string, phoneNumber string, password string) User {

	// Generate a random salt
	salt := make([]byte, 16)
	rand.Read(salt)

	// Check if the phoneNumber parameter is empty
	// This is necessary for weirdness with gorm and the unique constraint
	var phonePtr *string
	if phoneNumber != "" {
		phonePtr = &phoneNumber
	}

	// Hash the password with the salt
	hasedPass, _ := shared.HashPassword(password, salt)

	return User{
		BaseModel: BaseModel{
			ID: uuid.New(),
		},
		Username:    username,
		Email:       email,
		PhoneNumber: phonePtr,
		Password:    *hasedPass,
		Salt:        salt,
		Privacy:     shared.PrivacyPublic,
		Role:        shared.RoleUser,
	}
}

func (ur *UserRepo) GetByEmail(email string) (*User, error) {
	var user User
	result := ur.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (ur *UserRepo) GetByUsername(username string) (*User, error) {
	var user User
	result := ur.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (ur *UserRepo) UsernameExists(username string) bool {
	var user User
	result := ur.db.Where("username = ?", username).First(&user)
	return !result.RecordNotFound()
}

func (ur *UserRepo) EmailExists(email string) bool {
	var user User
	result := ur.db.Where("email = ?", email).First(&user)
	return !result.RecordNotFound()
}

func (ur *UserRepo) PhoneNumberExists(phoneNumber string) bool {
	var user User
	result := ur.db.Where("phone_number = ?", phoneNumber).First(&user)
	return !result.RecordNotFound()
}

// IsVerifiedByID checks if a user with the given ID has verified their email or phone number.
func (ur *UserRepo) IsVerifiedByID(id *uuid.UUID) (bool, error) {
	user, err := ur.GetByID(id)
	if err != nil {
		return false, err
	}

	return user.VerifiedEmail || user.VerifiedPhoneNumber, nil
}

func (ur *UserRepo) CheckCredentials(username, password string) (*User, error) {
	user, err := ur.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if !user.VerifyPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
