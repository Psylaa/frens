package service

import (
	"github.com/bwoff11/frens/pkg/config"
	"github.com/bwoff11/frens/pkg/database"
)

type Service struct {
	Auth     *AuthService
	Block    *BlockService
	Bookmark *BookmarkService
	Feed     *FeedService
	Follow   *FollowService
	Like     *LikeService
	Media    *MediaService
	Post     *PostService
	User     *UserService
}

func New(db *database.Database, config *config.APIConfig) *Service {
	return &Service{
		Auth: &AuthService{
			Database:    db,
			JWTSecret:   []byte(config.TokenSecret),
			JWTDuration: config.TokenDuration,
		},
		Block:    &BlockService{Database: db},
		Bookmark: &BookmarkService{Database: db},
		Feed:     &FeedService{Database: db},
		Follow:   &FollowService{Database: db},
		Like:     &LikeService{Database: db},
		Media:    &MediaService{Database: db},
		Post:     &PostService{Database: db},
		User:     &UserService{Database: db},
	}
}
