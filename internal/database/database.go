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
	Posts     interface{ Base[Post] }
	Users     interface{ Base[User] }
	Blocks    interface{ Interactor[Block] }
	Bookmarks interface{ Interactor[Bookmark] }
	Follows   interface{ Interactor[Follow] }
	Likes     interface{ Interactor[Like] }
}

type Block struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target User `gorm:"foreignKey:UserID"`
}

type Bookmark struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target Post `gorm:"foreignKey:PostID"`
}

type Follow struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target User `gorm:"foreignKey:UserID"`
}

type Like struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target Post `gorm:"foreignKey:PostID"`
}

func New(cfg *config.Config) (*Database, error) {
	logger.Log.Info().
		Str("host", cfg.Database.Host).
		Str("port", cfg.Database.Port).
		Str("user", cfg.Database.User).
		Str("password", cfg.Database.Password).
		Str("dbname", cfg.Database.DBName).
		Str("sslmode", cfg.Database.SSLMode).
		Msg("Connecting to Postgres database")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.Password, cfg.Database.SSLMode)

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	logger.Log.Info().Msg("Successfully connected to Postgres database")

	return initializeDatabase(db, cfg.Database.LogMode, cfg.Database.MaxIdleConns, cfg.Database.MaxOpenConns)
}

func initializeDatabase(db *gorm.DB, logMode bool, maxIdleConns int, maxOpenConns int) (*Database, error) {
	db.LogMode(logMode)

	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	db.AutoMigrate(&User{}, &Post{}, &Like{}, &Follow{}, &Block{}, &Bookmark{})
	logger.Log.Info().Msg("Auto migration completed")

	return &Database{
		Posts:     NewBaseRepo[Post](db),
		Users:     NewUserRepo(db),
		Blocks:    NewInteractorRepo[Block](db),
		Bookmarks: NewInteractorRepo[Bookmark](db),
		Follows:   NewInteractorRepo[Follow](db),
		Likes:     NewInteractorRepo[Like](db),
	}, nil
}
