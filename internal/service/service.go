package service

import (
	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
)

type Service struct {
	Bookmarks *BookmarkRepo
	Files     *FilesRepo
	Posts     *PostRepo
	Users     *UserRepo
}

var db *database.Database
var cfg *config.Config

func New(database *database.Database, configuration *config.Config) *Service {
	db = database
	return &Service{
		Bookmarks: &BookmarkRepo{},
	}
}
