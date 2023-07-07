package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Likes interface {
	Base[Like]
}

type Like struct {
	BaseModel
}

type LikeRepo struct {
	*BaseRepo[Like]
}

func NewLikeRepo(db *gorm.DB) Likes {
	return &LikeRepo{NewBaseRepo[Like](db)}
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
