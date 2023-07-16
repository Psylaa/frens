package database

import (
	"time"

	"github.com/bwoff11/frens/pkg/config"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB represents a connection pool to the database.
type Database struct {
	Conn  *gorm.DB
	Repos *Repositories
}

// Repositories represents the collections of Repo.
type Repositories struct {
	Block     *BlockRepo
	Bookmarks *BookmarkRepo
	Follows   *FollowRepo
	Likes     *LikeRepo
	Media     *MediaRepo
	Posts     *PostRepo
	Users     *UserRepo
}

type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
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

	db.Conn.Model(&Block{}).AddUniqueIndex("idx_block_user_blocked", "user_id", "blocked_id")
	db.Conn.Model(&Bookmark{}).AddUniqueIndex("idx_bookmark_user_post", "user_id", "post_id")
	db.Conn.Model(&Follow{}).AddUniqueIndex("idx_follow_user_followed", "user_id", "followed_id")
	db.Conn.Model(&Like{}).AddUniqueIndex("idx_like_user_post", "user_id", "post_id")

	if config.DevMode {
		db.Conn.DropTableIfExists(
			&Block{},
			&Bookmark{},
			&Follow{},
			&Like{},
			&Media{},
			&Post{},
			&User{})
	}

	db.Conn.AutoMigrate(
		&Block{},
		&Bookmark{},
		&Follow{},
		&Like{},
		&Media{},
		&Post{},
		&User{},
	)

	db.Repos = &Repositories{
		Block:     &BlockRepo{Conn: db.Conn},
		Bookmarks: &BookmarkRepo{Conn: db.Conn},
		Follows:   &FollowRepo{Conn: db.Conn},
		Likes:     &LikeRepo{Conn: db.Conn},
		Media:     &MediaRepo{Conn: db.Conn},
		Posts:     &PostRepo{Conn: db.Conn},
		Users:     &UserRepo{Conn: db.Conn},
	}

	return &db, nil
}
