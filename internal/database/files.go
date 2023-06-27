package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type File struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Extension string    `json:"extension"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Owner     uuid.UUID `gorm:"type:uuid" json:"owner"`
}

func CreateFile(owner uuid.UUID, extension string) (*File, error) {
	file := &File{
		ID:        uuid.New(),
		Extension: extension,
		Owner:     owner,
	}
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
