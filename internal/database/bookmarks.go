package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	BaseModel
	UserID   uuid.UUID `json:"userId"`
	StatusID uuid.UUID `json:"statusId"`
}

func GetBookmarks(statusID uuid.UUID) ([]Bookmark, error) {
	var bookmarks []Bookmark
	if err := db.Where("status_id = ?", statusID).Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func CreateBookmark(userID, statusID uuid.UUID) (*Bookmark, error) {
	newBookmark := Bookmark{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    userID,
		StatusID:  statusID,
	}

	if err := db.Create(&newBookmark).Error; err != nil {
		return nil, err
	}

	return &newBookmark, nil
}

func DeleteBookmark(userID, statusID uuid.UUID) error {
	var bookmark Bookmark
	if err := db.Where("user_id = ? AND status_id = ?", userID, statusID).First(&bookmark).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("bookmark does not exist")
		}
		return err
	}

	if err := db.Delete(&bookmark).Error; err != nil {
		return err
	}

	return nil
}

func HasUserBookmarked(userID, statusID uuid.UUID) (bool, error) {
	var count int
	if err := db.Model(&Bookmark{}).Where("user_id = ? AND status_id = ?", userID, statusID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
