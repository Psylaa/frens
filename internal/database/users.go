package database

import (
	"crypto/rand"

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
}

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Salt     []byte

	Bio *string `gorm:"size:1024"`

	Privacy shared.Privacy
	Role    shared.Role

	Avatar   File `gorm:"foreignKey:AvatarID"`
	Cover    File `gorm:"foreignKey:CoverID"`
	CoverID  *uuid.UUID
	AvatarID *uuid.UUID
}

type UserRepo struct {
	*BaseRepo[User]
}

func NewUserRepo(db *gorm.DB) Users {
	return &UserRepo{NewBaseRepo[User](db)}
}

func NewUser(username string, email string, password string) User {

	// Generate a random salt
	salt := make([]byte, 16)
	rand.Read(salt)

	// Hash the password with the salt
	hasedPass, _ := shared.HashPassword(password, salt)

	return User{
		BaseModel: BaseModel{
			ID: uuid.New(),
		},
		Username: username,
		Email:    email,
		Password: *hasedPass,
		Salt:     salt,
		Privacy:  shared.PrivacyPublic,
		Role:     shared.RoleUser,
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
