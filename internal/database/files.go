package database

import (
	"log"
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
	PostID    uuid.UUID `gorm:"type:uuid" json:"postId"` // Used for many to one - dont delete
}

type FileRepo struct {
	db *gorm.DB
}

func (fr *FileRepo) CreateFile(owner uuid.UUID, extension string) (*File, error) {
	file := &File{
		ID:        uuid.New(),
		Extension: extension,
		Owner:     owner,
	}
	if err := fr.db.Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (fr *FileRepo) GetFile(id uuid.UUID) (*File, error) {
	log.Println("GetFile")
	var file File
	if err := fr.db.Where("id = ?", id).First(&file).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &file, nil
}

func (fr *FileRepo) GetFiles(ids []uuid.UUID) ([]File, error) {
	var files []File
	if err := fr.db.Where("id IN (?)", ids).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr *FileRepo) GetFilesByOwner(owner uuid.UUID) ([]File, error) {
	var files []File
	if err := fr.db.Where("owner = ?", owner).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr *FileRepo) UpdateFile(file *File) error {
	return fr.db.Save(file).Error
}

func (fr *FileRepo) DeleteFile(id uuid.UUID) error {
	file, err := fr.GetFile(id)
	if err != nil {
		return err
	}
	if file == nil {
		return gorm.ErrRecordNotFound
	}
	return fr.db.Delete(file).Error
}
