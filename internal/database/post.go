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
	CreateWithMedia(post *Post, mediaIDs []*uuid.UUID, fr Files) error
	GetByID(id *uuid.UUID, requestorID *uuid.UUID) (*Post, error)
	GetByIDs(ids []*uuid.UUID, requestorID *uuid.UUID) ([]*Post, error)
	GetByUserIDs(userIDs []*uuid.UUID, cursor time.Time, count int, requestorID *uuid.UUID) ([]*Post, error)
}

// Post struct represents the post table in the database with appropriate gorm tags.
type Post struct {
	BaseModel
	AuthorID     uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"` // UUID of the post author
	Author       User           `gorm:"foreignKey:AuthorID" json:"author"`   // The author of the post
	Privacy      shared.Privacy `gorm:"type:varchar(20)" json:"privacy"`     // Privacy setting of the post
	Text         string         `gorm:"type:text" json:"text"`               // Text content of the post
	Media        []*File        `json:"media"`                               // Media content of the post
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

func (pr *PostRepo) CreateWithMedia(post *Post, mediaIDs []*uuid.UUID, fr Files) error {
	logger.DebugLogRequestReceived("database", "PostRepo", "CreateWithMedia")

	tx := pr.db.Begin()

	// Create post
	err := tx.Create(post).Error
	if err != nil {
		logger.Log.Error().Err(err).Msg("error creating post")
		tx.Rollback()
		return err
	}

	// Update the files
	for _, mediaID := range mediaIDs {
		err := fr.UpdatePostIDInTx(tx, *mediaID, post.ID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("error updating post id in file")
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetByID returns the post with the given ID, preloading the Author data
func (pr *PostRepo) GetByID(id *uuid.UUID, requestorID *uuid.UUID) (*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByID")

	var post Post
	err := pr.db.
		Preload("Author").
		Preload("Media").
		Where("id = ?", id).
		First(&post).Error

	if err != nil {
		return nil, err
	}

	likes, err := pr.getLikes([]uuid.UUID{*id}, *requestorID)
	if err != nil {
		return nil, err
	}

	bookmarks, err := pr.getBookmarks([]uuid.UUID{*id}, *requestorID)
	if err != nil {
		return nil, err
	}

	post.IsLiked = likes[*id]
	post.IsBookmarked = bookmarks[*id]

	return &post, nil
}

func (pr *PostRepo) GetByIDs(ids []*uuid.UUID, requestorID *uuid.UUID) ([]*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByIDs")

	var posts []*Post
	err := pr.db.
		Preload("Author").
		Preload("Media").
		Where("id IN (?)", ids).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	postIDs := make([]uuid.UUID, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}

	likes, err := pr.getLikes(postIDs, *requestorID)
	if err != nil {
		return nil, err
	}

	bookmarks, err := pr.getBookmarks(postIDs, *requestorID)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		post.IsLiked = likes[post.ID]
		post.IsBookmarked = bookmarks[post.ID]
	}

	return posts, nil
}

func (pr *PostRepo) GetByUserIDs(userIDs []*uuid.UUID, cursor time.Time, count int, requestorID *uuid.UUID) ([]*Post, error) {
	logger.DebugLogRequestReceived("database", "PostRepo", "GetByUserIDs")

	var posts []*Post
	result := pr.db.
		Preload("Author").
		Preload("Media").
		Where("author_id IN (?) AND created_at < ?", userIDs, cursor).
		Order("created_at DESC").
		Limit(count).
		Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	// I wasnt able to get the more complex queries working to get this running in one request, so we need to loop through the posts and get isliked and isbookmarked for each post
	// I dont think DB performance is a huge issue at this point, but this is something to keep in mind for the future
	for _, post := range posts {
		var like Like
		result = pr.db.Where("user_id = ? AND post_id = ?", post.AuthorID, post.ID).First(&like)
		if result.Error == nil {
			post.IsLiked = true
		}

		var bookmark Bookmark
		result = pr.db.Where("user_id = ? AND post_id = ?", post.AuthorID, post.ID).First(&bookmark)
		if result.Error == nil {
			post.IsBookmarked = true
		}
	}

	return posts, nil
}

// getLikes returns a map of postID to isLiked status.
func (pr *PostRepo) getLikes(postIDs []uuid.UUID, userID uuid.UUID) (map[uuid.UUID]bool, error) {
	var likes []Like
	err := pr.db.
		Where("post_id IN (?) AND user_id = ?", postIDs, userID).
		Find(&likes).Error

	if err != nil {
		return nil, err
	}

	isLiked := make(map[uuid.UUID]bool)
	for _, like := range likes {
		isLiked[*like.PostID] = true
	}
	return isLiked, nil
}

// getBookmarks returns a map of postID to isBookmarked status.
func (pr *PostRepo) getBookmarks(postIDs []uuid.UUID, userID uuid.UUID) (map[uuid.UUID]bool, error) {
	var bookmarks []Bookmark
	err := pr.db.
		Where("post_id IN (?) AND user_id = ?", postIDs, userID).
		Find(&bookmarks).Error

	if err != nil {
		return nil, err
	}

	isBookmarked := make(map[uuid.UUID]bool)
	for _, bookmark := range bookmarks {
		isBookmarked[bookmark.PostID] = true
	}
	return isBookmarked, nil
}
