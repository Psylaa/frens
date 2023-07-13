package database

import (
	"time"

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
	ReadBySource(sourceID uuid.UUID, count *int, cursor *time.Time) ([]T, error)
	ReadByTarget(targetID uuid.UUID, count *int, cursor *time.Time) ([]T, error)
	DeleteBySource(sourceID uuid.UUID) error
	DeleteByTarget(targetID uuid.UUID) error
}

func NewInteractorRepo[T Entity](db *gorm.DB) Interactor[T] {
	return &InteractorRepo[T]{NewBaseRepo[T](db)}
}

// ReadBySource reads all interactors by source user id.
func (r *InteractorRepo[T]) ReadBySource(sourceID uuid.UUID, count *int, cursor *time.Time) ([]T, error) {
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
func (r *InteractorRepo[T]) ReadByTarget(targetID uuid.UUID, count *int, cursor *time.Time) ([]T, error) {
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

// DeleteBySource deletes all interactors by source user id.
func (r *InteractorRepo[T]) DeleteBySource(sourceID uuid.UUID) error {
	var interactor T
	return r.db.Where("source_id = ?", sourceID).Delete(&interactor).Error
}

// DeleteByTarget deletes all interactors by target user id.
func (r *InteractorRepo[T]) DeleteByTarget(targetID uuid.UUID) error {
	var interactor T
	return r.db.Where("target_id = ?", targetID).Delete(&interactor).Error
}
