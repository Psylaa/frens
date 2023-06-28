package database

import (
	"errors"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Follow struct {
	BaseModel
	SourceID uuid.UUID `json:"sourceId"`
	TargetID uuid.UUID `json:"targetId"`
}

type FollowRepo struct {
	db *gorm.DB
}

func (fr *FollowRepo) CreateFollow(sourceID uuid.UUID, targetID uuid.UUID) (*Follow, error) {
	newFollow := Follow{
		BaseModel: BaseModel{ID: uuid.New()},
		SourceID:  sourceID,
		TargetID:  targetID,
	}

	if err := fr.db.Create(&newFollow).Error; err != nil {
		logger.Log.Error().Msgf("Failed to create follow: %v", err)
		return nil, err
	}

	logger.Log.Debug().Msgf("Created follow: %v", newFollow)
	return &newFollow, nil
}

func (fr *FollowRepo) DeleteFollow(sourceID, targetID uuid.UUID) error {
	var follow Follow
	if err := fr.db.Where("source_id = ? AND target_id = ?", sourceID, targetID).First(&follow).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		logger.Log.Warn().Msgf("Follow not found: %v", err)
		return err
	}

	if err := fr.db.Delete(&follow).Error; err != nil {
		logger.Log.Error().Msgf("Failed to delete follow: %v", err)
		return err
	}

	logger.Log.Debug().Msgf("Deleted follow: %v", follow)
	return nil
}

func (fr *FollowRepo) GetFollowers(targetID uuid.UUID) ([]Follow, error) {
	var follows []Follow
	if err := fr.db.Where("target_id = ?", targetID).Find(&follows).Error; err != nil {
		logger.Log.Error().Msgf("Failed to get followers: %v", err)
		return nil, err
	}

	return follows, nil
}

func (fr *FollowRepo) GetFollowing(sourceID uuid.UUID) ([]Follow, error) {
	var following []Follow
	if err := fr.db.Where("source_id = ?", sourceID).Find(&following).Error; err != nil {
		logger.Log.Error().Msgf("Failed to get following: %v", err)
		return nil, err
	}

	return following, nil
}

func (fr *FollowRepo) DoesFollowExist(sourceID, targetID uuid.UUID) (bool, error) {
	var follow Follow
	if err := fr.db.Where("source_id = ? AND target_id = ?", sourceID, targetID).First(&follow).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		logger.Log.Error().Msgf("Failed to check if follow exists: %v", err)
		return false, err
	}

	return true, nil
}
