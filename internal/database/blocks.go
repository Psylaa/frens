package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Blocks interface {
	Base[Block]
	GetByID(id *uuid.UUID) (*Block, error)
	GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]Block, error)
}

type Block struct {
	BaseModel
	SourceUserID uuid.UUID
	TargetUserID uuid.UUID
}

type BlockRepo struct {
	*BaseRepo[Block]
}

func NewBlockRepo(db *gorm.DB) Blocks {
	return &BlockRepo{NewBaseRepo[Block](db)}
}

func (br *BlockRepo) GetByID(id *uuid.UUID) (*Block, error) {
	var block Block
	result := br.db.First(&block, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &block, nil
}

// Returns a paginated list of entities. Count defines the number of entities, and offset the starting position.
func (br *BlockRepo) GetByUserID(userID *uuid.UUID, count *int, offset *int) ([]Block, error) {
	var blocks []Block
	result := br.db.Where("source_user_id = ?", userID).Limit(count).Offset(offset).Find(&blocks)
	if result.Error != nil {
		return nil, result.Error
	}

	return blocks, nil
}
