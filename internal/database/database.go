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
	Followers *FollowerRepo
	Likes     *LikeRepo
	//Media     *MediaRepo
	Posts *PostRepo
	Users *UserRepo
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

	db.AutoMigrate(&User{}, &Post{}, &Like{}, &Follower{}, &Bookmark{}, &File{})
	logger.Log.Info().Msg("Auto migration completed")

	// Manually create the composite unique index
	db.Model(&Like{}).AddUniqueIndex("idx_user_status", "user_id", "status_id")
	logger.Log.Info().Msg("Created unique index for Like")

	return &Database{
		DB:        db,
		Bookmarks: &BookmarkRepo{db: db},
		Files:     &FileRepo{db: db},
		Followers: &FollowerRepo{db: db},
		Likes:     &LikeRepo{db: db},
		Posts:     &PostRepo{db: db},
		Users:     &UserRepo{db: db},
	}, nil
}
