package service

import (
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
)

var JWTSigningKey []byte
var JWTDuration time.Duration
var defaultBio string

type Service struct {
	Blocks    *BlockRepo
	Bookmarks *BookmarkRepo
	Feed      *FeedRepo
	Likes     *LikeRepo
	Posts     *PostRepo
	Users     *UserRepo
}

func New(configuration *config.Config) *Service {

	// Store config values
	JWTSigningKey = []byte(configuration.Server.JWTSecret)
	JWTDuration = time.Duration(configuration.Server.JWTDuration) * time.Hour
	defaultBio = configuration.Users.DefaultBio

	database, err := database.New(configuration)
	if err != nil {
		panic(err)
	}

	s := &Service{
		Blocks:    &BlockRepo{Database: database},
		Bookmarks: &BookmarkRepo{Database: database},
		Feed:      &FeedRepo{Database: database},
		Likes:     &LikeRepo{Database: database},
		Posts:     &PostRepo{Database: database},
		Users:     &UserRepo{Database: database},
	}

	return s
}
