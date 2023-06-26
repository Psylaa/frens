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
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Type  FileType  `gorm:"type:varchar(100)"`
	Owner uuid.UUID `gorm:"type:uuid"`
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
