package database

import (
	"log"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	BaseModel
	UserID   uuid.UUID `json:"userId"`
	StatusID uuid.UUID `json:"statusId"`
	Owner    User      `gorm:"foreignKey:UserID" json:"owner"`
}

type BookmarkRepo struct {
	db *gorm.DB
}

func (br *BookmarkRepo) GetBookmarkByID(bookmarkID *uuid.UUID) (*Bookmark, error) {
	var bookmark Bookmark
	if err := br.db.
		Preload("Owner").
		Where("id = ?", bookmarkID).
		First(&bookmark).
		Error; err != nil {
		return nil, err
	}

	log.Println(bookmark.Owner)
	log.Println(bookmark.Owner)
	log.Println(bookmark.Owner)

	return &bookmark, nil
}

func (br *BookmarkRepo) GetBookmarksByIDs(statusID uuid.UUID) ([]*Bookmark, error) {
	var bookmarks []*Bookmark
	if err := br.db.
		Preload("Owner").
		Where("status_id = ?", statusID).
		Find(&bookmarks).
		Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (br *BookmarkRepo) GetBookmarkCount(statusID uuid.UUID) (int, error) {
	var count int
	if err := br.db.
		Model(&Bookmark{}).
		Where("status_id = ?", statusID).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (br *BookmarkRepo) CreateBookmark(userID, statusID uuid.UUID) (*Bookmark, error) {
	newBookmark := &Bookmark{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    userID,
		StatusID:  statusID,
	}

	if err := br.db.Create(newBookmark).Error; err != nil {
		return nil, err
	}

	return newBookmark, nil
}

func (br *BookmarkRepo) DeleteBookmark(userID, statusID uuid.UUID) error {
	var bookmark Bookmark
	if err := br.db.Where("user_id = ? AND status_id = ?", userID, statusID).First(&bookmark).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := br.db.Delete(&bookmark).Error; err != nil {
		return err
	}

	return nil
}

func (br *BookmarkRepo) HasUserBookmarked(userID, statusID uuid.UUID) (bool, error) {
	var count int
	if err := br.db.Model(&Bookmark{}).Where("user_id = ? AND status_id = ?", userID, statusID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
