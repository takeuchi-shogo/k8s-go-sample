package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type BlockRepository struct{}

func (repo *BlockRepository) FindByID(db *gorm.DB, id int) (*models.Blocks, error) {
	block := &models.Blocks{}
	if err := db.First(block, id).Error; !errors.Is(err, nil) {
		return &models.Blocks{}, fmt.Errorf("%w", err)
	}
	return block, nil
}

func (repo *BlockRepository) Create(db *gorm.DB, b *models.Blocks) (*models.Blocks, error) {
	block := &models.Blocks{}
	if err := db.Create(block).Error; !errors.Is(err, nil) {
		return &models.Blocks{}, fmt.Errorf("%w", err)
	}
	return block, nil
}
