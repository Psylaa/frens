package database

import (
	"errors"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/jinzhu/gorm"
)

type Posts interface {
	Base[models.Post]
}

type PostsRepo struct {
	*BaseRepo[models.Post]
}

func NewPostRepo(db *gorm.DB) Posts {
	if db == nil {
		logger.Error(logger.LogMessage{
			Package:  "database",
			Function: "NewPostRepo",
			Message:  "Attempted to create new post repo with nil database",
		}, errors.New("database is nil"))
	}

	return &PostsRepo{NewBaseRepo[models.Post](db)}
}
