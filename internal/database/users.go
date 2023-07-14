package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Verified bool   `gorm:"default:false"`
}

type Users interface {
	Base[User]
}

type UserRepo struct {
	*BaseRepo[User]
}

func NewUserRepo(db *gorm.DB) Users {
	return &UserRepo{NewBaseRepo[User](db)}
}

// Override. Since read by user applied to all but the actual user
// We need to translate the id to the actual user id
func (r *UserRepo) ReadByUser(id uuid.UUID) ([]User, error) {
	var users []User
	err := r.db.Where("id = ?", id).Find(&users).Error
	return users, err
}

// Override. Same as above.
func (r *UserRepo) ReadByUsers(ids []uuid.UUID) ([]User, error) {
	var users []User
	err := r.db.Where("id IN (?)", ids).Find(&users).Error
	return users, err
}
