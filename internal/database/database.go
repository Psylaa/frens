package database

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Database struct {
	*gorm.DB
	Bookmarks *BookmarkRepo
	Files     *FileRepo
	Follows   *FollowRepo
	Likes     *LikeRepo
	Posts     *PostRepo
	Users     *UserRepo
}

/*
	The interfaces are here just to make it easier to organize the code.
	They are not necessary, but they do make it easier to see what methods
	are available for each model and make sure we have proper naming
	standards, consistency, scope, and coverage.

	Naming convention:
	- Except when accepting a pointer to a struct, name should be "By" + type
*/

type BookmarkRepoInterface interface {
	Create(bookmark *Bookmark) error
	DeleteByID(bookmarkID *uuid.UUID) error
	ExistsByPostAndUserID(postID *uuid.UUID, userID *uuid.UUID) (bool, error)
	GetByID(bookmarkID *uuid.UUID) (*Bookmark, error)
	GetByPostID(postID *uuid.UUID, count *int, offest *int) ([]*Bookmark, error)
	GetByPostAndUserID(userID *uuid.UUID, postID *uuid.UUID) (*Bookmark, error)
	GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]*Bookmark, error)
}

var _ BookmarkRepoInterface = &BookmarkRepo{} // Ensure interface is implemented

type FileRepoInterface interface {
	Create(file *File) error
	DeleteByID(fileID *uuid.UUID) error
	GetByID(fileID *uuid.UUID) (*File, error)
	GetByPostID(postID *uuid.UUID) ([]*File, error)
	GetByUserID(userID *uuid.UUID) ([]*File, error)
}

var _ FileRepoInterface = &FileRepo{} // Ensure interface is implemented

type FollowRepoInterface interface {
	Create(follow *Follow) error
	DeleteBySourceAndTargetID(sourceId uuid.UUID, targetId uuid.UUID) error
	ExistsBySourceAndTargetID(sourceId uuid.UUID, targetId uuid.UUID) (bool, error)
	GetByID(followId *uuid.UUID) (*Follow, error)
	GetBySourceID(sourceId uuid.UUID, count *int, offset *int) ([]*Follow, error)
	GetByTargetID(targetId uuid.UUID, count *int, offset *int) ([]*Follow, error)
}

type LikeRepoInterface interface {
	GetByIDs(likeIds []*uuid.UUID) ([]*Like, error)
	GetByPostID(postID *uuid.UUID) ([]*Like, error)
	GetByUserID(userID *uuid.UUID) ([]*Like, error)

	GetCount_ByPostID(postID *uuid.UUID) (int, error)
	GetCount_ByUserID(userID *uuid.UUID) (int, error)

	Create(like *Like) (*Like, error)
	Delete(userId *uuid.UUID, postId *uuid.UUID) error
	Exists(userId *uuid.UUID, postId *uuid.UUID) (bool, error)
}

type PostRepoInterface interface {
	GetByIDs(postIDs []*uuid.UUID) ([]*Post, error)
	GetByUserID(userID *uuid.UUID) ([]*Post, error)

	GetCount_ByUserID(userID *uuid.UUID) (int, error)

	Create(post *Post) (*Post, error)
	Delete(postID *uuid.UUID) error
}

type UserRepoInterface interface {
	GetByIDs(userID []*uuid.UUID) ([]*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)

	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(userID *uuid.UUID) error
}

// New initializes a new database connection
func New(cfg *config.Config) (*Database, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.Password, cfg.Database.SSLMode)

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	db.LogMode(cfg.Database.LogMode)
	logger.Log.Info().Msg("Successfully connected to database")

	db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)

	db.AutoMigrate(&User{}, &Post{}, &Like{}, &Follow{}, &Bookmark{}, &File{})
	logger.Log.Info().Msg("Auto migration completed")

	// Manually create the composite unique index
	//db.Model(&Like{}).AddUniqueIndex("idx_user_post", "user_id", "post_id")
	logger.Log.Info().Msg("Created unique index for Like")

	return &Database{
		DB:        db,
		Bookmarks: &BookmarkRepo{db: db},
		Files:     &FileRepo{db: db},
		Follows:   &FollowRepo{db: db},
		Likes:     &LikeRepo{db: db},
		Posts:     &PostRepo{db: db, Likes: &LikeRepo{db: db}, Bookmarks: &BookmarkRepo{db: db}},
		Users:     &UserRepo{db: db, Follows: &FollowRepo{db: db}},
	}, nil
}
