package database

import (
	"fmt"

	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	Conn *gorm.DB
}

func New(config *config.DatabaseConfig) (*Database, error) {
	var db Database
	var err error

	dbInfo := "host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.User +
		" dbname=" + config.Name +
		" password=" + config.Password +
		" sslmode=" + config.SSLMode

	db.Conn, err = gorm.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	db.Conn.LogMode(config.LogMode)

	if config.DevMode {
		db.Conn.DropTableIfExists(&models.Block{}, &models.Bookmark{}, &models.Follow{}, &models.Like{}, &models.Media{}, &models.Post{}, &models.User{})
	}

	db.Conn.AutoMigrate(&models.Block{}, &models.Bookmark{}, &models.Follow{}, &models.Like{}, &models.Media{}, &models.Post{}, &models.User{})

	err = db.Conn.Model(&models.Block{}).AddUniqueIndex("idx_block_user_blocked", "user_id", "blocked_id").Error
	if err != nil {
		return nil, fmt.Errorf("failed to add unique index for Block: %v", err)
	}

	err = db.Conn.Model(&models.Bookmark{}).AddUniqueIndex("idx_bookmark_user_post", "user_id", "post_id").Error
	if err != nil {
		return nil, fmt.Errorf("failed to add unique index for Bookmark: %v", err)
	}

	err = db.Conn.Model(&models.Follow{}).AddUniqueIndex("idx_follow_user_followed", "user_id", "followed_id").Error
	if err != nil {
		return nil, fmt.Errorf("failed to add unique index for Follow: %v", err)
	}

	err = db.Conn.Model(&models.Like{}).AddUniqueIndex("idx_like_user_post", "user_id", "post_id").Error
	if err != nil {
		return nil, fmt.Errorf("failed to add unique index for Like: %v", err)
	}

	return &db, nil
}
