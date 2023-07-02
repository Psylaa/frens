package service

import (
	"log"

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
	cfg = configuration

	if db == nil {
		log.Panic("Database pointer provided to service package is nil")
	}

	if cfg == nil {
		log.Panic("Config pointer provided to service package is nil")
	}

	return &Service{
		Bookmarks: &BookmarkRepo{},
	}
}
