package database

import (
	"github.com/bwoff11/frens/internal/models"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	*BaseRepo[models.User]
}

type Users interface {
	Base[models.User]
	ReadByEmail(email string) (models.User, error)
}

func NewUserRepo(db *gorm.DB) Users {
	return &UserRepo{NewBaseRepo[models.User](db)}
}

func (r *UserRepo) ReadByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
