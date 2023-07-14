package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository struct{ db *gorm.DB }

type UserRepo interface {
	Create(user *models.User) error
	Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (r *UserRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.User, error) {
	var users []models.User
	query := r.db

	if limit != nil {
		query = query.Limit(*limit)
	}

	if cursor != nil {
		query = query.Where("created_at < ?", cursor)
	}

	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	err := query.Find(&users).Error
	return users, err
}

func (ur *UserRepository) ReadByEmail(email string) (models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(id uuid.UUID) error {
	return ur.db.Where("id = ?", id).Delete(&models.User{}).Error
}
