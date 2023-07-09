package database

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Entity interface{}

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Base represents a generic CRUD interface for a database entity
type Base[T Entity] interface {
	// Create inserts a new entity into the database
	Create(entity *T) error

	// Update modifies an existing entity in the database
	Update(entity *T) error

	// DeleteByID deletes an entity by its ID
	Delete(entity *T) error
}

type BaseRepo[T Entity] struct {
	db *gorm.DB
}

// Returns a new BaseRepo instance
func NewBaseRepo[T Entity](db *gorm.DB) *BaseRepo[T] {
	return &BaseRepo[T]{db: db}
}

// Creates an entity in the database
func (repo *BaseRepo[T]) Create(entity *T) error {
	logger.DebugLogRequestReceived("database", "BaseRepo", "Create")
	result := repo.db.Create(entity)
	return result.Error
}

// Updates an entity in the database
func (repo *BaseRepo[T]) Update(entity *T) error {
	logger.DebugLogRequestReceived("database", "BaseRepo", "Update")
	result := repo.db.Save(entity)
	return result.Error
}

// Deletes an entity with the given ID
func (repo *BaseRepo[T]) Delete(entity *T) error {
	logger.DebugLogRequestReceived("database", "BaseRepo", "Delete")
	result := repo.db.Delete(entity)
	return result.Error
}
