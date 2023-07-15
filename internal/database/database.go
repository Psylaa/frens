package database

import (
	"fmt"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	db        *gorm.DB
	Block     *BlockRepository
	Bookmarks *BookmarkRepository
	Follows   *FollowRepository
	Likes     *LikeRepository
	Media     *MediaRepository
	Posts     *PostRepository
	Users     *UserRepository
}

func New(cfg *config.Config) (*Database, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.Password, cfg.Database.SSLMode)

	gormDB, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		logger.Fatal(logger.LogMessage{
			Package:  "database",
			Function: "New",
			Message:  "Failed to connect to database: " + err.Error(),
		}, err)
		return nil, err
	}

	gormDB.LogMode(cfg.Database.LogMode)
	gormDB.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)
	gormDB.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)

	gormDB.AutoMigrate(
		&models.Block{},
		&models.Bookmark{},
		&models.Follow{},
		&models.Like{},
		&models.Media{},
		&models.Post{},
		&models.Post{},
		&models.User{},
	)

	return &Database{
		db:        gormDB,
		Block:     &BlockRepository{db: gormDB},
		Bookmarks: &BookmarkRepository{db: gormDB},
		Follows:   &FollowRepository{db: gormDB},
		Likes:     &LikeRepository{db: gormDB},
		Media:     &MediaRepository{db: gormDB},
		Posts:     &PostRepository{db: gormDB},
		Users:     &UserRepository{db: gormDB},
	}, nil
}
