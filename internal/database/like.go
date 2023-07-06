package database

import "github.com/jinzhu/gorm"

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
