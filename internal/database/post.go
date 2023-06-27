package database

import (
	"time"

	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

// Post represents a post update by a user.
type Post struct {
	BaseModel
	OwnerID uuid.UUID      `gorm:"type:uuid" json:"owner"`
	Privacy shared.Privacy `json:"privacy"`
	Text    string         `json:"text"`
	Media   []Media        `json:"media" gorm:"ForeignKey:PostID"`
}

// GetPost gets a post update by ID.
func GetPost(id uuid.UUID) (*Post, error) {
	var post Post
	if err := db.Preload("Media").First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPostesByUserID gets all post updates by a user.
func GetPostsByUserID(userID uuid.UUID) ([]Post, error) {
	var postes []Post
	if err := db.Preload("Media").
		Order("created_at desc").
		Find(&postes, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return postes, nil
}

// GetPostesByUserIDs gets a limited number of post updates from multiple users,
// older than the provided timestamp.
func GetPostsByUserIDs(userIDs []uuid.UUID, cursor time.Time, limit int) ([]Post, error) {
	var postes []Post
	if err := db.Preload("Media").
		Where("user_id IN (?) AND created_at < ?", userIDs, cursor).
		Order("created_at desc").
		Limit(limit).
		Find(&postes).Error; err != nil {
		return nil, err
	}
	return postes, nil
}

// CreatePost creates a new post update.
func CreatePost(post *Post) error {
	err := db.Create(post).Error
	return err
}

// DeletePost deletes a post update by ID.
func DeletePost(postID uuid.UUID) error {
	err := db.Delete(&Post{}, "id = ?", postID).Error
	return err
}
