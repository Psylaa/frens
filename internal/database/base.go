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

	// GetByID fetches an entity by its ID
	GetByID(id *uuid.UUID) (*T, error)

	// DeleteByID deletes an entity by its ID
	DeleteByID(id *uuid.UUID) error

	// GetPaginated returns a paginated list of entities
	GetPaginated(count, offset *int) ([]*T, error)
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

// Returns an entity with the given ID
func (repo *BaseRepo[T]) GetByID(id *uuid.UUID) (*T, error) {
	logger.DebugLogRequestReceived("database", "BaseRepo", "GetByID")
	var entity T
	result := repo.db.Where("id = ?", id).First(&entity)
	return &entity, result.Error
}

// Deletes an entity with the given ID
func (repo *BaseRepo[T]) DeleteByID(id *uuid.UUID) error {
	logger.DebugLogRequestReceived("database", "BaseRepo", "DeleteByID")
	var entity T
	result := repo.db.Where("id = ?", id).Delete(&entity)
	return result.Error
}

// Returns a paginated list of entities. Count defines the number of entities, and offset the starting position.
func (repo *BaseRepo[T]) GetPaginated(count, offset *int) ([]*T, error) {
	logger.DebugLogRequestReceived("database", "BaseRepo", "GetPaginated")
	var entities []*T

	db := repo.db

	if count != nil {
		db = db.Limit(*count)
	}

	if offset != nil {
		db = db.Offset(*offset)
	}

	result := db.Find(&entities)

	return entities, result.Error
}
