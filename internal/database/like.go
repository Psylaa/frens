package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Likes interface {
	Base[Like]
	GetByPostIDAndUserID(postID *uuid.UUID, userID *uuid.UUID) (*Like, error)
	Exists(userID *uuid.UUID, postID *uuid.UUID) (bool, error)
}

type Like struct {
	BaseModel
	UserID *uuid.UUID `gorm:"type:uuid;not null"`
	User   User       `gorm:"foreignKey:UserID"`
	PostID *uuid.UUID `gorm:"type:uuid;not null"`
	Post   Post       `gorm:"foreignKey:PostID"`
}

type LikeRepo struct {
	*BaseRepo[Like]
}

func NewLikeRepo(db *gorm.DB) Likes {
	return &LikeRepo{NewBaseRepo[Like](db)}
}

func (lr *LikeRepo) GetByPostIDAndUserID(postID *uuid.UUID, userID *uuid.UUID) (*Like, error) {
	logger.DebugLogRequestReceived("database", "LikeRepo", "GetByUserIDAndPostID")

	var like Like
	result := lr.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&like)
	if result.Error != nil {
		return nil, result.Error
	}

	return &like, nil
}

func (lr *LikeRepo) Exists(userID *uuid.UUID, postID *uuid.UUID) (bool, error) {
	logger.DebugLogRequestReceived("database", "LikeRepo", "Exists")

	var like Like
	result := lr.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&like)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}
