package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Like struct {
	BaseModel
	UserID   uuid.UUID `json:"userId"`
	StatusID uuid.UUID `json:"statusId"`
}

func CreateLike(userID, statusID uuid.UUID) (*Like, error) {
	newLike := Like{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    userID,
		StatusID:  statusID,
	}

	if err := db.Create(&newLike).Error; err != nil {
		return nil, err
	}

	return &newLike, nil
}

func DeleteLike(userID, statusID uuid.UUID) error {
	var like Like
	if err := db.Where("user_id = ? AND status_id = ?", userID, statusID).First(&like).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("like does not exist")
		}
		return err
	}

	if err := db.Delete(&like).Error; err != nil {
		return err
	}

	return nil
}
