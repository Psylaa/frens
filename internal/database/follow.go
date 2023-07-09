package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Follows interface {
	Base[Follow]
	GetFollowing(userID *uuid.UUID) ([]*User, error)
}

type Follow struct {
	BaseModel
	SourceID uuid.UUID `gorm:"type:uuid;not null"`
	TargetID uuid.UUID `gorm:"type:uuid;not null"`
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

func (fr *FollowRepo) GetFollowing(sourceID *uuid.UUID) ([]*User, error) {
	logger.DebugLogRequestReceived("database", "FollowRepo", "GetFollowing")

	var users []*User
	result := fr.db.Model(&User{}).
		Joins("JOIN follows ON follows.target_id = users.id").
		Where("follows.source_id = ?", sourceID).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
