package database

import "github.com/jinzhu/gorm"

type Bookmarks interface {
	Base[Bookmark]
}

type Bookmark struct {
	BaseModel
}

type BookmarkRepo struct {
	*BaseRepo[Bookmark]
}

func NewBookmarkRepo(db *gorm.DB) Bookmarks {
	return &BookmarkRepo{NewBaseRepo[Bookmark](db)}
}
