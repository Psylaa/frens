package database

import (
	"fmt"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
	Blocks    Blocks
	Bookmarks Bookmarks
	Files     Files
	Follows   Follows
	Likes     Likes
	Posts     Posts
	Users     Users
}

// New initializes a new database connection
func New(cfg *config.Config) (*Database, error) {
	switch cfg.Database.Type {
	case config.SQLite:
		return NewSQLite(cfg)
	case config.Postgres:
		return NewPostgres(cfg)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}
}

func NewSQLite(cfg *config.Config) (*Database, error) {
	logger.Log.Info().
		Str("path", cfg.Database.SQLite.Path).
		Msg("Connecting to SQLite database")

	db, err := gorm.Open("sqlite3", cfg.Database.SQLite.Path)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to connect to SQLite database")
		return nil, err
	}

	logger.Log.Info().Msg("Successfully connected to SQLite database")

	return initializeDatabase(db, cfg.Database.SQLite.LogMode, cfg.Database.SQLite.MaxIdleConns, cfg.Database.SQLite.MaxOpenConns)
}

func NewPostgres(cfg *config.Config) (*Database, error) {
	logger.Log.Info().
		Str("host", cfg.Database.Postgres.Host).
		Str("port", cfg.Database.Postgres.Port).
		Str("user", cfg.Database.Postgres.User).
		Str("password", cfg.Database.Postgres.Password).
		Str("dbname", cfg.Database.Postgres.DBName).
		Str("sslmode", cfg.Database.Postgres.SSLMode).
		Msg("Connecting to Postgres database")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Postgres.Host, cfg.Database.Postgres.Port, cfg.Database.Postgres.User, cfg.Database.Postgres.DBName, cfg.Database.Postgres.Password, cfg.Database.Postgres.SSLMode)

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	logger.Log.Info().Msg("Successfully connected to Postgres database")

	return initializeDatabase(db, cfg.Database.Postgres.LogMode, cfg.Database.Postgres.MaxIdleConns, cfg.Database.Postgres.MaxOpenConns)
}

func initializeDatabase(db *gorm.DB, logMode bool, maxIdleConns int, maxOpenConns int) (*Database, error) {
	db.LogMode(logMode)

	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	db.AutoMigrate(&User{}, &Post{}, &Like{}, &Follow{}, &Bookmark{}, &File{})
	logger.Log.Info().Msg("Auto migration completed")

	// Manually create the composite unique index
	db.Model(&Like{}).AddUniqueIndex("idx_like_user_post", "user_id", "post_id")
	db.Model(&Bookmark{}).AddUniqueIndex("idx_bm_user_post", "user_id", "post_id")

	return &Database{
		Blocks:    NewBlockRepo(db),
		Bookmarks: NewBookmarkRepo(db),
		Files:     NewFileRepo(db),
		Follows:   NewFollowRepo(db),
		Likes:     NewLikeRepo(db),
		Posts:     NewPostRepo(db),
		Users:     NewUserRepo(db),
	}, nil
}
