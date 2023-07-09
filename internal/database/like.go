package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Likes interface {
	Create(like *Like) error
	GetByID(id *uuid.UUID) (*Like, error)
	Exists(userID *uuid.UUID, postID *uuid.UUID) (bool, error)
}

type Like struct {
	BaseModel
	UserID *uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
	PostID *uuid.UUID
	Post   Post `gorm:"foreignKey:PostID"`
}

type LikeRepo struct {
	*BaseRepo[Like]
}

func NewLikeRepo(db *gorm.DB) Likes {
	return &LikeRepo{NewBaseRepo[Like](db)}
}

func (lr *LikeRepo) Create(like *Like) error {
	logger.DebugLogRequestReceived("database", "LikeRepo", "Create")

	result := lr.db.Create(like)
	return result.Error
}

func (lr *LikeRepo) GetByID(id *uuid.UUID) (*Like, error) {
	logger.DebugLogRequestReceived("database", "LikeRepo", "GetByID")

	var like Like
	result := lr.db.First(&like, id)
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
