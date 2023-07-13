package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Entity interface{}

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Base[T Entity] interface {
	Create(entity *T) error
	Read(id uuid.UUID) (*T, error)
	ReadMany(ids []uuid.UUID) ([]T, error)
	Update(entity *T) error
	Delete(id uuid.UUID) error
}

type BaseRepo[T Entity] struct {
	db *gorm.DB
}

func NewBaseRepo[T Entity](db *gorm.DB) *BaseRepo[T] {
	return &BaseRepo[T]{db: db}
}

// Creates a new entity
func (repo *BaseRepo[T]) Create(entity *T) error {
	return repo.db.Create(entity).Error
}

// Reads an entity by id
func (repo *BaseRepo[T]) Read(id uuid.UUID) (*T, error) {
	var entity T
	err := repo.db.Where("id = ?", id).First(&entity).Error
	return &entity, err
}

// Reads many entities by ids
func (repo *BaseRepo[T]) ReadMany(ids []uuid.UUID) ([]T, error) {
	var entities []T
	err := repo.db.Where("id IN (?)", ids).Find(&entities).Error
	return entities, err
}

// Updates an entity
func (repo *BaseRepo[T]) Update(entity *T) error {
	return repo.db.Save(entity).Error
}

// Deletes an entity
func (repo *BaseRepo[T]) Delete(id uuid.UUID) error {
	var entity T
	return repo.db.Where("id = ?", id).Delete(&entity).Error
}
