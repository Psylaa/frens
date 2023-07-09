package database

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Posts represents the interface for a Post repository
type Posts interface {
	Base[Post]
	GetByID(id *uuid.UUID, requestorID *uuid.UUID) (*Post, error)
	GetByUserIDs(userIDs []*uuid.UUID, cursor time.Time, count int, requestorID *uuid.UUID) ([]*Post, error)
}

// Post struct represents the post table in the database with appropriate gorm tags.
type Post struct {
	BaseModel
	AuthorID     uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"` // UUID of the post author
	Author       User           `gorm:"foreignKey:AuthorID" json:"author"`   // The author of the post
	Privacy      shared.Privacy `gorm:"type:varchar(20)" json:"privacy"`     // Privacy setting of the post
	Text         string         `gorm:"type:text" json:"text"`               // Text content of the post
	Media        []*File        `gorm:"-" json:"media"`                      // Media content of the post
	MediaIDs     []*uuid.UUID   `gorm:"-" json:"-"`                          // Helper field to hold the Media ID's while processing a request
	IsLiked      bool           `gorm:"-" json:"isLiked"`                    // Indicates if post is liked by user
	IsBookmarked bool           `gorm:"-" json:"isBookmarked"`               // Indicates if post is bookmarked by user
}

// PostRepo struct represents the Post repository
type PostRepo struct {
	*BaseRepo[Post]
}

// NewPostRepo initializes and returns a Post repository
func NewPostRepo(db *gorm.DB) Posts {
	return &PostRepo{NewBaseRepo[Post](db)}
}

// GetByID returns the post with the given ID, preloading the Author and Media data
func (pr *PostRepo) GetByID(id *uuid.UUID, requestorID *uuid.UUID) (*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByID")

	var post Post
	result := pr.db.
		Preload("Author").
		Select("posts.*, CASE WHEN likes.id IS NOT NULL THEN true ELSE false END AS is_liked").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", requestorID).
		Joins("LEFT JOIN bookmarks ON bookmarks.post_id = posts.id AND bookmarks.user_id = ?", requestorID).
		Where("id = ?", id).
		First(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	// Manual preload for Media
	var media []*File
	result = pr.db.Where("id IN (?)", post.MediaIDs).Find(&media)
	if result.Error != nil {
		return nil, result.Error
	}
	post.Media = media

	return &post, nil
}

func (pr *PostRepo) GetByUserIDs(userIDs []*uuid.UUID, cursor time.Time, count int, requestorID *uuid.UUID) ([]*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByUserIDs")

	var posts []*Post
	result := pr.db.
		Preload("Author").
		Select("posts.*, CASE WHEN likes.id IS NOT NULL THEN true ELSE false END AS is_liked, CASE WHEN bookmarks.id IS NOT NULL THEN true ELSE false END AS is_bookmarked").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", requestorID).
		Joins("LEFT JOIN bookmarks ON bookmarks.post_id = posts.id AND bookmarks.user_id = ?", requestorID).
		Where("posts.author_id IN (?) AND posts.created_at < ?", userIDs, cursor).
		Order("posts.created_at DESC").
		Limit(count).
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	// Manual preload for Media
	for _, post := range posts {
		var media []*File
		result = pr.db.Where("id IN (?)", post.MediaIDs).Find(&media)
		if result.Error != nil {
			return nil, result.Error
		}
		post.Media = media
	}

	return posts, nil
}
