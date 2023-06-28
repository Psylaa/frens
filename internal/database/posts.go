package database

import (
	"time"

	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Post represents a post update by a user.
type Post struct {
	BaseModel
	OwnerID uuid.UUID      `json:"ownerId"`
	Privacy shared.Privacy `json:"privacy"`
	Text    string         `json:"text"`
}

// PostRepo provides access to the Post storage.
type PostRepo struct {
	db *gorm.DB
}

func (pr *PostRepo) GetPost(id uuid.UUID) (*Post, error) {
	var post Post
	if err := pr.db.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (pr *PostRepo) GetPostsByUserID(userID uuid.UUID) ([]Post, error) {
	var posts []Post
	if err := pr.db.
		Order("created_at desc").
		Find(&posts, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *PostRepo) GetPostsByUserIDs(userIDs []uuid.UUID, cursor time.Time, limit int) ([]Post, error) {
	var posts []Post
	if err := pr.db.
		Where("owner_id IN (?) AND created_at < ?", userIDs, cursor).
		Order("created_at desc").
		Limit(limit).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *PostRepo) CreatePost(ownerID uuid.UUID, text string, privacy shared.Privacy) (*Post, error) {
	post := &Post{
		BaseModel: BaseModel{ID: uuid.New()},
		OwnerID:   ownerID,
		Privacy:   privacy,
		Text:      text,
	}
	if err := pr.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (pr *PostRepo) DeletePost(postID uuid.UUID) error {
	err := pr.db.Delete(&Post{}, "id = ?", postID).Error
	return err
}

func (pr *PostRepo) GetLatestPublicPosts(cursor time.Time, limit int) ([]*Post, error) {
	var posts []*Post
	err := pr.db.
		Where("created_at <= ? AND privacy = ?", cursor, "PUBLIC").
		Order("created_at desc").
		Limit(limit).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
