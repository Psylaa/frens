package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type FileType string

const (
	ProfilePicture FileType = "ProfilePicture"
	UserBanner     FileType = "UserBanner"
	StatusImage    FileType = "StatusImage"
	StatusVideo    FileType = "StatusVideo"
	StatusFile     FileType = "StatusFile"
)

type File struct {
	BaseModel
	Type  string    `gorm:"type:varchar(100)"`
	Owner uuid.UUID `gorm:"type:uuid"`
	ID    uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"` // Duplicated from BaseModel to make it easier to use
}

func CreateFile(file *File) (*File, error) {
	if err := db.Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func GetFile(id uuid.UUID) (*File, error) {
	var file File
	if err := db.Where("id = ?", id).First(&file).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

func UpdateFile(file *File) error {
	return db.Save(file).Error
}

func DeleteFile(id uuid.UUID) error {
	file, err := GetFile(id)
	if err != nil {
		return err
	}
	if file == nil {
		return gorm.ErrRecordNotFound
	}
	return db.Delete(file).Error
}
