package database

import (
	"errors"
	"time"

	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Post represents a post update by a user.
type Post struct {
	BaseModel
	OwnerID uuid.UUID      `json:"owner"`
	Privacy shared.Privacy `json:"privacy"`
	Text    string         `json:"text"`
	Media   []Media        `json:"media" gorm:"ForeignKey:PostID"`
}

// PostRepo provides access to the Post storage.
type PostRepo struct {
	db *gorm.DB
}

func (pr *PostRepo) GetPost(id uuid.UUID) (*Post, error) {
	var post Post
	if err := pr.db.Preload("Media").First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (pr *PostRepo) GetPostsByUserID(userID uuid.UUID) ([]Post, error) {
	var posts []Post
	if err := pr.db.Preload("Media").
		Order("created_at desc").
		Find(&posts, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *PostRepo) GetPostsByUserIDs(userIDs []uuid.UUID, cursor time.Time, limit int) ([]Post, error) {
	var posts []Post
	if err := pr.db.Preload("Media").
		Where("user_id IN (?) AND created_at < ?", userIDs, cursor).
		Order("created_at desc").
		Limit(limit).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *PostRepo) CreatePost(post *Post) error {
	if post == nil {
		return errors.New("provided post is not valid")
	}
	err := pr.db.Create(post).Error
	return err
}

func (pr *PostRepo) DeletePost(postID uuid.UUID) error {
	err := pr.db.Delete(&Post{}, "id = ?", postID).Error
	return err
}
