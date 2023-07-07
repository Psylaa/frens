package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

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

func (fr *FileRepo) GetByID(id *uuid.UUID) (*File, error) {
	logger.DebugLogRequestReceived("database", "FileRepo", "GetByID")

	var file File
	result := fr.db.First(&file, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &file, nil
}
