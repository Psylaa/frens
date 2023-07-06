package database

import (
	"fmt"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
	Bookmarks Bookmarks
	Files     Files
	Follows   Follows
	Likes     Likes
	Posts     Posts
	Users     Users
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
		Bookmarks: NewBookmarkRepo(db),
		Files:     NewFileRepo(db),
		Follows:   NewFollowRepo(db),
		Likes:     NewLikeRepo(db),
		Posts:     NewPostRepo(db),
		Users:     NewUserRepo(db),
	}, nil
}
