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
	*gorm.DB
	Posts Posts
	Users Users
	//Blocks
	//Bookmarks
	//Follows
	//Likes
}

type Block struct {
	InteractorModel
	Source models.User `gorm:"foreignKey:UserID"`
	Target models.User `gorm:"foreignKey:UserID"`
}

type Bookmark struct {
	InteractorModel
	Source models.User `gorm:"foreignKey:UserID"`
	Target Post        `gorm:"foreignKey:PostID"`
}

type Follow struct {
	InteractorModel
	Source models.User `gorm:"foreignKey:UserID"`
	Target models.User `gorm:"foreignKey:UserID"`
}

type Like struct {
	InteractorModel
	Source models.User `gorm:"foreignKey:UserID"`
	Target Post        `gorm:"foreignKey:PostID"`
}

func New(cfg *config.Config) (*Database, error) {
	logger.Info(logger.LogMessage{
		Package:  "database",
		Function: "New",
		Message:  "Initializing database",
	})

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.Password, cfg.Database.SSLMode)

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	return initializeDatabase(db, cfg.Database.LogMode, cfg.Database.MaxIdleConns, cfg.Database.MaxOpenConns)
}

func initializeDatabase(db *gorm.DB, logMode bool, maxIdleConns int, maxOpenConns int) (*Database, error) {
	if db == nil {
		logger.Fatal(logger.LogMessage{
			Package:  "database",
			Function: "initializeDatabase",
			Message:  "Attempted to initialize a nil database",
		}, nil)
	}

	db.LogMode(logMode)

	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	db.AutoMigrate(&models.User{}, &Post{}, &Like{}, &Follow{}, &Block{}, &Bookmark{})

	newDB := &Database{
		DB:    db,
		Posts: NewPostRepo(db),
		Users: NewUserRepo(db),
		//Blocks: NewBlocksRepo(db),
		//Bookmarks: NewBookmarksRepo(db),
		//Follows:   NewFollowsRepo(db),
		//Likes:     NewLikesRepo(db),
	}

	logger.Info(logger.LogMessage{
		Package:  "database",
		Function: "initializeDatabase",
		Message:  "Database initialized",
	})

	return newDB, nil
}
