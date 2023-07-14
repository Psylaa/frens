package service

import (
	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
)

type Service struct {
	Blocks    *BlockRepo
	Bookmarks *BookmarkRepo
	Feed      *FeedRepo
	Likes     *LikeRepo
	Posts     *PostRepo
	Users     *UserRepo
}

func New(configuration *config.Config) *Service {

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
