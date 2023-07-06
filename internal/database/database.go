package database

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	SourceID  uuid.UUID `gorm:"type:uuid;"`
	TargetID  uuid.UUID `gorm:"type:uuid;"`
}

type Database struct {
	*gorm.DB
	Bookmarks interface{ Base[Bookmark] }
	Files     interface{ Base[File] }
	Follows   interface{ Base[Follow] }
	Likes     interface{ Base[Like] }
	Posts     interface{ Base[Post] }
	Users     Users
}

type Bookmark struct {
	BaseModel
	UserID uuid.UUID
	PostID uuid.UUID
	Owner  User `gorm:"foreignKey:UserID"`
}

type File struct {
	BaseModel
	Extension string
}

type Post struct {
	BaseModel
	Author   User `gorm:"foreignKey:AuthorID"`
	AuthorID uuid.UUID
	Privacy  shared.Privacy
	Text     string
	Media    []*File `gorm:"foreignKey:PostID;AssociationForeignKey:ID" json:"media"`
}

type User struct {
	BaseModel
	Username    string `gorm:"unique"`
	Email       string
	Bio         string
	Password    string
	Avatar      File `gorm:"foreignKey:AvatarID"`
	AvatarID    uuid.UUID
	Cover       File `gorm:"foreignKey:CoverID"`
	CoverID     uuid.UUID
	Privacy     shared.Privacy
	Role        shared.Role
	IsFollowing bool `gorm:"-"`
}

type Follow struct {
	BaseModel
}

type Like struct {
	BaseModel
}

type Entity interface{}

// Base represents a generic CRUD interface for a database entity
type Base[T Entity] interface {
	// Create inserts a new entity into the database
	Create(entity *T) error

	// Update modifies an existing entity in the database
	Update(entity *T) error

	// GetByID fetches entities by their ID
	GetByID(id *uuid.UUID) ([]*T, error)

	// DeleteByID deletes an entity by its ID
	DeleteByID(id *uuid.UUID) error

	// GetBySourceID fetches entities by their SourceID
	GetBySourceID(sourceID *uuid.UUID, limit *int, offset *int) ([]*T, error)

	// GetByTargetID fetches entities by their TargetID
	GetByTargetID(targetID *uuid.UUID, limit *int, offset *int) ([]*T, error)

	// DeleteBySourceAndTargetID deletes an entity by its SourceID and TargetID
	DeleteBySourceAndTargetID(entity *T, sourceID *uuid.UUID, targetID *uuid.UUID) error

	// ExistsBySourceAndTargetID checks if an entity with the given SourceID and TargetID exists
	ExistsBySourceAndTargetID(entity *T, sourceID *uuid.UUID, targetID *uuid.UUID) (bool, error)
}

type Users interface {
	Base[User]
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
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
		Bookmarks: NewBaseRepo[Bookmark](db),
		Files:     NewBaseRepo[File](db),
		Follows:   NewBaseRepo[Follow](db),
		Likes:     NewBaseRepo[Like](db),
		Posts:     NewBaseRepo[Post](db),
		//Users:     NewUserRepo(db), // This needs to be defined separately as it has additional methods
	}, nil
}
