package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Like struct {
	BaseModel
	UserID uuid.UUID
	PostID uuid.UUID
}

type LikeRepo struct {
	db *gorm.DB
}

func (lr *LikeRepo) GetByPostID(postID uuid.UUID) ([]*Like, error) {
	var likes []*Like
	if err := lr.db.Where("post_id = ?", postID).Find(&likes).Error; err != nil {
		return nil, err
	}

	return likes, nil
}

func (lr *LikeRepo) GetCountByPostID(postID uuid.UUID) (*int, error) {
	var count int
	if err := lr.db.Model(&Like{}).Where("post_id = ?", postID).Count(&count).Error; err != nil {
		return nil, err
	}
	return &count, nil
}

func (lr *LikeRepo) Create(userID *uuid.UUID, postID *uuid.UUID) (*Like, error) {
	newLike := &Like{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    *userID,
		PostID:    *postID,
	}

	if err := lr.db.Create(newLike).Error; err != nil {
		return nil, err
	}

	return newLike, nil
}

func (lr *LikeRepo) Delete(userID *uuid.UUID, postID *uuid.UUID) error {
	var like Like
	if err := lr.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error; err != nil {
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

func (lr *LikeRepo) Exists(userID *uuid.UUID, postID *uuid.UUID) (bool, error) {
	var count int
	if err := lr.db.Model(&Like{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
