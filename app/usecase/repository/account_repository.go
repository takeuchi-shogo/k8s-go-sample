package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindByID(db *gorm.DB, id int) (*models.Accounts, error)
	FirstByEmail(db *gorm.DB, email string) (*models.Accounts, error)
	Create(db *gorm.DB, a *models.Accounts) (*models.Accounts, error)
	Save(db *gorm.DB, a *models.Accounts) (*models.Accounts, error)
}
