package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BaseRepo[T Entity] struct {
	db *gorm.DB
}

// Returns a new BaseRepo instance
func NewBaseRepo[T Entity](db *gorm.DB) *BaseRepo[T] {
	return &BaseRepo[T]{db: db}
}

// Creates an entity in the database
func (repo *BaseRepo[T]) Create(entity *T) error {
	result := repo.db.Create(entity)
	return result.Error
}

// Updates an entity in the database
func (repo *BaseRepo[T]) Update(entity *T) error {
	result := repo.db.Save(entity)
	return result.Error
}

// Returns an entity with the given ID
func (repo *BaseRepo[T]) GetByID(id *uuid.UUID) ([]*T, error) {
	var entities []*T
	result := repo.db.Where("id = ?", id).Find(&entities)
	return entities, result.Error
}

// Deletes an entity with the given ID
func (repo *BaseRepo[T]) DeleteByID(id *uuid.UUID) error {
	var entity T
	result := repo.db.Where("id = ?", id).Delete(entity)
	return result.Error
}

// Returns entities with the given IDs
func (repo *BaseRepo[T]) GetBySourceID(sourceID *uuid.UUID, limit *int, offset *int) ([]*T, error) {
	var entities []*T
	result := repo.db.Where("source_id = ?", sourceID).Limit(limit).Offset(offset).Find(&entities)
	return entities, result.Error
}

// Returns entities with the given IDs
func (repo *BaseRepo[T]) GetByTargetID(targetID *uuid.UUID, limit *int, offset *int) ([]*T, error) {
	var entities []*T
	result := repo.db.Where("target_id = ?", targetID).Limit(limit).Offset(offset).Find(&entities)
	return entities, result.Error
}

// Returns entities with the given IDs
func (repo *BaseRepo[T]) GetBySourceAndTargetID(sourceID *uuid.UUID, targetID *uuid.UUID) (*T, error) {
	var entity T
	result := repo.db.First(&entity, "source_id = ? AND target_id = ?", sourceID, targetID)
	return &entity, result.Error
}

// Checks if an entity with the given SourceID and TargetID exists
func (repo *BaseRepo[T]) ExistsBySourceAndTargetID(entity *T, sourceID *uuid.UUID, targetID *uuid.UUID) (bool, error) {
	var count int64
	result := repo.db.Model(entity).Where("source_id = ? AND target_id = ?", sourceID, targetID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

// Deletes an entity with the given SourceID and TargetID
func (repo *BaseRepo[T]) DeleteBySourceAndTargetID(entity *T, sourceID *uuid.UUID, targetID *uuid.UUID) error {
	result := repo.db.Delete(entity, "source_id = ? AND target_id = ?", sourceID, targetID)
	return result.Error
}
