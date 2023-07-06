package database

import (
	"log"
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type File struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	PostID    uuid.UUID `gorm:"type:uuid"` // Used for many to one - dont delete
}

type FileRepo struct {
	db *gorm.DB
}

func (fr *FileRepo) Create(file *File) error {
	logger.DebugLogRequestReceived("database", "files", "Create")
	if err := fr.db.Create(file).Error; err != nil {
		return err
	}
	return nil
}

func (fr *FileRepo) DeleteByID(fileID *uuid.UUID) error {
	log.Println("DeleteFile")
	if err := fr.db.Delete(&File{}, "id = ?", fileID).Error; err != nil {
		return err
	}
	return nil
}

func (fr *FileRepo) GetByID(fileID *uuid.UUID) (*File, error) {
	log.Println("GetFile")
	var file File
	if err := fr.db.Where("id = ?", fileID).First(&file).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &file, nil
}

func (fr *FileRepo) GetByPostID(postID *uuid.UUID) ([]*File, error) {
	var files []*File
	if err := fr.db.Where("post_id = ?", postID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr *FileRepo) GetByUserID(userID *uuid.UUID) ([]*File, error) {
	var files []*File
	if err := fr.db.Where("owner = ?", userID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
