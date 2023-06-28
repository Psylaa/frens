package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Like struct {
	BaseModel
	UserID   uuid.UUID `json:"userId"`
	StatusID uuid.UUID `json:"statusId"`
}

type LikeRepo struct {
	db *gorm.DB
}

func (lr *LikeRepo) GetLikes(statusID uuid.UUID) ([]Like, error) {
	var likes []Like
	if err := lr.db.Where("status_id = ?", statusID).Find(&likes).Error; err != nil {
		return nil, err
	}

	return likes, nil
}

func (lr *LikeRepo) CreateLike(userID, statusID uuid.UUID) (*Like, error) {
	newLike := Like{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    userID,
		StatusID:  statusID,
	}

	if err := lr.db.Create(&newLike).Error; err != nil {
		return nil, err
	}

	return &newLike, nil
}

func (lr *LikeRepo) DeleteLike(userID, statusID uuid.UUID) error {
	var like Like
	if err := lr.db.Where("user_id = ? AND status_id = ?", userID, statusID).First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := lr.db.Delete(&like).Error; err != nil {
		return err
	}

	return nil
}

func (lr *LikeRepo) HasUserLiked(userID, statusID uuid.UUID) (bool, error) {
	var count int
	if err := lr.db.Model(&Like{}).Where("user_id = ? AND status_id = ?", userID, statusID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
