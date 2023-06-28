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

type FollowerRepo struct {
	db *Database
}

func (fr *FollowerRepo) CreateFollower(followingID, followerID uuid.UUID) (*Follower, error) {
	newFollower := Follower{
		BaseModel:   BaseModel{ID: uuid.New()},
		FollowingID: followingID,
		FollowerID:  followerID,
	}

	if err := fr.db.DB.Create(&newFollower).Error; err != nil {
		return nil, err
	}

	return &newFollower, nil
}

func (fr *FollowerRepo) DeleteFollower(followingID, followerID uuid.UUID) error {
	var follower Follower
	if err := fr.db.DB.Where("following_id = ? AND follower_id = ?", followingID, followerID).First(&follower).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if err := fr.db.DB.Delete(&follower).Error; err != nil {
		return err
	}

	return nil
}

func (fr *FollowerRepo) GetFollowers(followingID uuid.UUID) ([]Follower, error) {
	var followers []Follower
	if err := fr.db.DB.Where("following_id = ?", followingID).Find(&followers).Error; err != nil {
		return nil, err
	}

	return followers, nil
}

func (fr *FollowerRepo) GetFollowing(followerID uuid.UUID) ([]Follower, error) {
	var following []Follower
	if err := fr.db.DB.Where("follower_id = ?", followerID).Find(&following).Error; err != nil {
		return nil, err
	}

	return following, nil
}
