package router

import (
	"github.com/bwoff11/frens/service/bookmark"
)

type BookmarksRepo struct {
	Service *bookmark.Service
}
