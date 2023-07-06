package database

import "github.com/jinzhu/gorm"

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
