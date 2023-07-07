package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Follows interface {
	Base[Follow]
}

type Follow struct {
	BaseModel
}

type FollowRepo struct {
	*BaseRepo[Follow]
}

func NewFollowRepo(db *gorm.DB) Follows {
	return &FollowRepo{NewBaseRepo[Follow](db)}
}

func (fr *FollowRepo) GetByID(id *uuid.UUID) (*Follow, error) {
	logger.DebugLogRequestReceived("database", "FollowRepo", "GetByID")

	var follow Follow
	result := fr.db.First(&follow, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &follow, nil
}
