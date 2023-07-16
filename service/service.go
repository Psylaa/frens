package service

import (
	"github.com/bwoff11/frens/pkg/database"
	"github.com/bwoff11/frens/service/block"
	"github.com/bwoff11/frens/service/bookmark"
	"github.com/bwoff11/frens/service/feed"
	"github.com/bwoff11/frens/service/follow"
	"github.com/bwoff11/frens/service/like"
	"github.com/bwoff11/frens/service/media"
	"github.com/bwoff11/frens/service/post"
	"github.com/bwoff11/frens/service/user"
)

type Service struct {
	Block    *block.Service
	Bookmark *bookmark.Service
	Feed     *feed.Service
	Follow   *follow.Service
	Like     *like.Service
	Media    *media.Service
	Post     *post.Service
	User     *user.Service
}

func New(db *database.Database) *Service {
	return &Service{
		Block:    block.New(db),
		Bookmark: bookmark.New(db),
		Follow:   follow.New(db),
		Like:     like.New(db),
		Media:    media.New(db),
		Post:     post.New(db),
		User:     user.New(db),
	}
}
