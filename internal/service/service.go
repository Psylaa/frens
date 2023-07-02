package service

import "github.com/bwoff11/frens/internal/database"

type Service struct {
	Bookmarks *BookmarkRepo
	Posts     *PostRepo
	Users     *UserRepo
}

var db *database.Database

func New(database *database.Database) *Service {
	db = database
	return &Service{
		Bookmarks: &BookmarkRepo{},
	}
}
