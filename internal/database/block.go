package database

import (
	"time"

	"github.com/bwoff11/frens/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BlockRepository struct{ db *gorm.DB }

func (b *BlockRepository) Create(block *models.Block) error {
	return b.db.Create(block).Error
}

func (b *BlockRepository) Read(limit *int, cursor *time.Time, ids ...uuid.UUID) ([]models.Block, error) {
	var blocks []models.Block
	query := b.db

	if limit != nil {
		query = query.Limit(*limit)
	}

	if cursor != nil {
		query = query.Where("created_at < ?", cursor)
	}

	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	err := query.Find(&blocks).Error
	return blocks, err
}

func (b *BlockRepository) Update(block *models.Block) error {
	return b.db.Save(block).Error
}

func (b *BlockRepository) Delete(id uuid.UUID) error {
	return b.db.Where("id = ?", id).Delete(&models.Block{}).Error
}

func (b *BlockRepository) GetBySourceID(id uuid.UUID) ([]models.Block, error) {
	var blocks []models.Block
	err := b.db.Where("source_id = ?", id).Find(&blocks).Error
	return blocks, err
}

func (b *BlockRepository) GetByTargetID(id uuid.UUID) ([]models.Block, error) {
	var blocks []models.Block
	err := b.db.Where("target_id = ?", id).Find(&blocks).Error
	return blocks, err
}

func (b *BlockRepository) DeleteBySourceID(id uuid.UUID) error {
	return b.db.Where("source_id = ?", id).Delete(&models.Block{}).Error
}

func (b *BlockRepository) DeleteByTargetID(id uuid.UUID) error {
	return b.db.Where("target_id = ?", id).Delete(&models.Block{}).Error
}

func (b *BlockRepository) Exists(sourceID, targetID uuid.UUID) (bool, error) {
	var count int
	err := b.db.Model(&models.Block{}).Where("source_id = ? AND target_id = ?", sourceID, targetID).Count(&count).Error
	return count > 0, err
}
