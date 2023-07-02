package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	BaseModel
	UserID uuid.UUID `json:"userId"`
	PostID uuid.UUID `json:"postId"`
	Owner  User      `gorm:"foreignKey:UserID" json:"owner"`
}

type BookmarkRepo struct {
	db *gorm.DB
}

func (br *BookmarkRepo) GetByID(bookmarkID *uuid.UUID) (*Bookmark, error) {
	var bookmark Bookmark
	if err := br.db.
		Preload("Owner").
		Where("id = ?", bookmarkID).
		First(&bookmark).
		Error; err != nil {
		return nil, err
	}

	return &bookmark, nil
}

func (br *BookmarkRepo) GetByPostID(postID *uuid.UUID) ([]*Bookmark, error) {
	var bookmarks []*Bookmark
	if err := br.db.
		Preload("Owner").
		Where("post_id = ?", postID).
		Find(&bookmarks).
		Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (br *BookmarkRepo) GetCountByPostID(postID *uuid.UUID) (int, error) {
	var count int
	if err := br.db.
		Model(&Bookmark{}).
		Where("status_id = ?", postID).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (br *BookmarkRepo) GetCountByUserID(userID *uuid.UUID) (int, error) {
	var count int
	if err := br.db.
		Model(&Bookmark{}).
		Where("user_id = ?", userID).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (br *BookmarkRepo) Create(userID *uuid.UUID, postID *uuid.UUID) (*Bookmark, error) {
	newBookmark := &Bookmark{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    *userID,
		PostID:    *postID,
	}

	if err := br.db.Create(newBookmark).Error; err != nil {
		return nil, err
	}

	return newBookmark, nil
}

func (br *BookmarkRepo) DeleteByID(userID *uuid.UUID, postID *uuid.UUID) (*Bookmark, error) {
	var bookmark Bookmark
	if err := br.db.
		Where("user_id = ? AND status_id = ?", userID, postID).
		First(&bookmark).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := br.db.Delete(&bookmark).Error; err != nil {
		return nil, err
	}

	return &bookmark, nil
}
