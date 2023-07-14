package database

import (
	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Users interface {
	Base[models.User]
}

type UserRepo struct {
	*BaseRepo[models.User]
}

func NewUserRepo(db *gorm.DB) Users {
	return &UserRepo{NewBaseRepo[models.User](db)}
}

// Override. Since read by user applied to all but the actual user
// We need to translate the id to the actual user id
func (r *UserRepo) ReadByUser(id uuid.UUID) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("id = ?", id).Find(&users).Error
	return users, err
}

// Override. Same as above.
func (r *UserRepo) ReadByUsers(ids []uuid.UUID) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("id IN (?)", ids).Find(&users).Error
	return users, err
}
