package database

import "github.com/jinzhu/gorm"

type Files interface {
	Base[File]
}

type File struct {
	BaseModel
	Extension string
}

type FileRepo struct {
	*BaseRepo[File]
}

func NewFileRepo(db *gorm.DB) Files {
	return &FileRepo{NewBaseRepo[File](db)}
}
