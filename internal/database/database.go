// internal/database/database.go
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

var db *gorm.DB

func InitDB(cfg *config.Config) error {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.Password, cfg.Database.SSLMode)

	var err error
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to connect to database")
		return err
	}
	logger.Log.Info().Msg("Successfully connected to database")

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Status{})
	db.AutoMigrate(&Media{})
	db.AutoMigrate(&Like{})
	db.AutoMigrate(&Follower{})
	db.AutoMigrate(&Bookmark{})
	logger.Log.Info().Msg("Auto migration completed")

	// Manually create the composite unique index
	db.Model(&Like{}).AddUniqueIndex("idx_user_status", "user_id", "status_id")
	logger.Log.Info().Msg("Created unique index for Like")

	return nil
}
