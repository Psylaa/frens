package database

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	BaseModel
	Author       User `gorm:"foreignKey:AuthorID"`
	AuthorID     uuid.UUID
	Privacy      shared.Privacy
	Text         string
	Media        []*File `gorm:"foreignKey:PostID;AssociationForeignKey:ID" json:"media"`
	IsLiked      bool    `gorm:"-"` // This is a virtual field that is not stored in the database
	IsBookmarked bool    `gorm:"-"` // This is a virtual field that is not stored in the database
}

type PostRepo struct {
	db        *gorm.DB
	Likes     *LikeRepo
	Bookmarks *BookmarkRepo
}

func (pr *PostRepo) GetByID(requestorID *uuid.UUID, postID *uuid.UUID) (*Post, error) {
	var post Post
	if err := pr.db.
		Preload("Author").
		Preload("Author.Avatar").
		Preload("Author.Cover").
		Preload("Media").
		First(&post, "id = ?", postID).Error; err != nil {
		return nil, err
	}

	// Get is liked
	likeExists, err := pr.Likes.Exists(requestorID, postID)
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error checking if post is liked: %s", err.Error())
		return nil, err
	}

	// Get is bookmarked
	bookmarkExists, err := pr.Bookmarks.Exists(requestorID, postID)
	if err != nil {
		logger.Log.Debug().Str("package", "database").Msgf("Error checking if post is bookmarked: %s", err.Error())
		return nil, err
	}

	// Set virtual fields
	post.IsLiked = likeExists
	post.IsBookmarked = bookmarkExists

	return &post, nil
}

func (pr *PostRepo) GetByUserID(userID *uuid.UUID) ([]*Post, error) {
	var posts []*Post
	if err := pr.db.
		Preload("Author").
		Preload("Author.Avatar").
		Preload("Author.Cover").
		Preload("Media").
		Order("created_at desc").
		Find(&posts, "author_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *PostRepo) GetByUserIDs(userIDs []uuid.UUID, cursor time.Time, limit int) ([]*Post, error) {
	var posts []*Post
	if err := pr.db.
		Preload("Author").
		Preload("Author.Avatar").
		Preload("Author.Cover").
		Preload("Media").
		Where("author_id IN (?) AND created_at < ?", userIDs, cursor).
		Order("created_at desc").
		Limit(limit).
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *PostRepo) GetLatestPublic(limit int) ([]Post, error) {
	var posts []Post
	err := pr.db.
		Preload("Author").
		Preload("Author.Avatar").
		Preload("Author.Cover").
		Preload("Media").
		Where("privacy = ?", "public").
		Order("created_at desc").
		Limit(limit).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *PostRepo) Create(authorID uuid.UUID, text string, privacy shared.Privacy, media []*File) (*Post, error) {
	post := Post{
		BaseModel: BaseModel{ID: uuid.New()},
		AuthorID:  authorID,
		Privacy:   privacy,
		Text:      text,
		Media:     media,
	}
	if err := pr.db.Create(&post).Error; err != nil {
		logger.Log.Error().
			Str("package", "database").
			Msgf("error creating post: %v", err)
		return nil, err
	}
	logger.Log.Debug().
		Str("package", "database").
		Msgf("successfully created post: %v", post)

	return &post, nil
}

func (pr *PostRepo) Delete(postID uuid.UUID) error {
	err := pr.db.Delete(&Post{}, "id = ?", postID).Error
	if err != nil {
		logger.Log.Error().
			Str("package", "database").
			Msgf("error deleting post: %v", err)
	}
	return err
}
