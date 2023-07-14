package database

import (
	"errors"
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type InteractorModel struct {
	BaseModel
	SourceID uuid.UUID `gorm:"type:uuid;not null"`
	TargetID uuid.UUID `gorm:"type:uuid;not null"`
}

type InteractorRepo[T Entity] struct {
	*BaseRepo[T]
}

type Interactor[T Entity] interface {
	Base[T]
	ReadBySourceID(sourceID uuid.UUID, count *int, cursor *time.Time) ([]T, error)
	ReadByTargetID(targetID uuid.UUID, count *int, cursor *time.Time) ([]T, error)
	CountBySourceID(sourceID uuid.UUID) (int, error)
	CountByTargetID(targetID uuid.UUID) (int, error)
	DeleteBySourceID(sourceID uuid.UUID) error
	DeleteByTargetID(targetID uuid.UUID) error
}

func NewInteractorRepo[T Entity](db *gorm.DB) Interactor[T] {
	if db == nil {
		logger.Error(logger.LogMessage{
			Package:  "database",
			Function: "NewInteractorRepo",
			Message:  "Attempted to create new interactor repo with nil database",
		}, errors.New("database is nil"))
	}

	return &InteractorRepo[T]{
		NewBaseRepo[T](db),
	}
}

// ReadBySource reads all interactors by source user id.
func (r *InteractorRepo[T]) ReadBySourceID(sourceID uuid.UUID, count *int, cursor *time.Time) ([]T, error) {
	var interactors []T
	query := r.db.Where("source_id = ?", sourceID)
	if count != nil {
		query = query.Limit(*count)
	}
	if cursor != nil {
		query = query.Where("created_at < ?", *cursor)
	}
	err := query.Find(&interactors).Error
	return interactors, err
}

// ReadByTarget reads all interactors by target user id.
func (r *InteractorRepo[T]) ReadByTargetID(targetID uuid.UUID, count *int, cursor *time.Time) ([]T, error) {
	var interactors []T
	query := r.db.Where("target_id = ?", targetID)
	if count != nil {
		query = query.Limit(*count)
	}
	if cursor != nil {
		query = query.Where("created_at < ?", *cursor)
	}
	err := query.Find(&interactors).Error
	return interactors, err
}

// CountBySourceID counts all interactors by source user id.
func (r *InteractorRepo[T]) CountBySourceID(sourceID uuid.UUID) (int, error) {
	var entity T
	var count int
	err := r.db.Model(&entity).Where("source_id = ?", sourceID).Count(&count).Error
	return count, err
}

// CountByTargetID counts all interactors by target user id.
func (r *InteractorRepo[T]) CountByTargetID(targetID uuid.UUID) (int, error) {
	var entity T
	var count int
	err := r.db.Model(&entity).Where("target_id = ?", targetID).Count(&count).Error
	return count, err
}

// DeleteBySource deletes all interactors by source user id.
func (r *InteractorRepo[T]) DeleteBySourceID(sourceID uuid.UUID) error {
	var interactor T
	return r.db.Where("source_id = ?", sourceID).Delete(&interactor).Error
}

// DeleteByTarget deletes all interactors by target user id.
func (r *InteractorRepo[T]) DeleteByTargetID(targetID uuid.UUID) error {
	var interactor T
	return r.db.Where("target_id = ?", targetID).Delete(&interactor).Error
}
