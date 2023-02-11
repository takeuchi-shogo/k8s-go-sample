package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type BlockRepository interface {
	FindByID(db *gorm.DB, id int) (*models.Blocks, error)
	Create(db *gorm.DB, block *models.Blocks) (*models.Blocks, error)
}
