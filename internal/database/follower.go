package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Follower struct {
	BaseModel
	FollowingID uuid.UUID `json:"followingId"`
	FollowerID  uuid.UUID `json:"followerId"`
}

func CreateFollower(followingID, followerID uuid.UUID) (*Follower, error) {
	newFollower := Follower{
		BaseModel:   BaseModel{ID: uuid.New()},
		FollowingID: followingID,
		FollowerID:  followerID,
	}

	if err := db.Create(&newFollower).Error; err != nil {
		return nil, err
	}

	return &newFollower, nil
}

func DeleteFollower(followingID, followerID uuid.UUID) error {
	var follower Follower
	if err := db.Where("following_id = ? AND follower_id = ?", followingID, followerID).First(&follower).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if err := db.Delete(&follower).Error; err != nil {
		return err
	}

	return nil
}

func GetFollowers(followingID uuid.UUID) ([]Follower, error) {
	var followers []Follower
	if err := db.Where("following_id = ?", followingID).Find(&followers).Error; err != nil {
		return nil, err
	}

	return followers, nil
}
