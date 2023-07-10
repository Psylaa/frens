package database

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Files interface {
	Base[File]
	GetByID(id *uuid.UUID) (*File, error)
	UpdatePostIDInTx(tx *gorm.DB, fileID, postID uuid.UUID) error
}

type File struct {
	BaseModel
	OwnerID   uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	PostID    uuid.UUID `gorm:"type:uuid"`
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
	result := fr.db.First(&file, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &file, nil
}

func (fr *FileRepo) UpdatePostIDInTx(tx *gorm.DB, fileID, postID uuid.UUID) error {
	logger.DebugLogRequestReceived("database", "FileRepo", "UpdatePostIDInTx")

	err := tx.Model(&File{}).Where("id = ?", fileID).Update("post_id", postID).Error
	if err != nil {
		logger.Log.Error().Err(err).Msg("error updating post id in file")
	}

	return err
}
