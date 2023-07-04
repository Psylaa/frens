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

func (fr *FileRepo) GetManyByID(fileIDs []*uuid.UUID) ([]*File, error) {
	var files []*File
	if err := fr.db.Where("id IN (?)", fileIDs).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr *FileRepo) GetByUserID(userID *uuid.UUID) ([]File, error) {
	var files []File
	if err := fr.db.Where("owner = ?", userID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr *FileRepo) Create(file *File) (*File, error) {
	logger.DebugLogRequestReceived("database", "files", "Create")
	if err := fr.db.Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (fr *FileRepo) Update(file *File) error {
	return fr.db.Save(file).Error
}

func (fr *FileRepo) DeleteByFile(file *File) error {
	return fr.db.Delete(file).Error
}

func (fr *FileRepo) DeleteByID(fileID *uuid.UUID) error {
	return fr.db.Where("id = ?", fileID).Delete(&File{}).Error
}
