package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type AccountRepository struct{}

func (repo *AccountRepository) FindByID(db *gorm.DB, id int) (*models.Accounts, error) {
	account := &models.Accounts{}
	if err := db.First(account, id).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed find account by id : %w", err)
	}
	return account, nil
}

func (repo *AccountRepository) Create(db *gorm.DB, a *models.Accounts) (*models.Accounts, error) {
	account := &models.Accounts{}
	account = a
	if err := db.Create(account).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed account create: %w", err)
	}
	return account, nil
}

func (repo *AccountRepository) Save(db *gorm.DB, a *models.Accounts) (*models.Accounts, error) {
	account := &models.Accounts{}
	account = a
	if err := db.Save(account).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed account create: %w", err)
	}
	return account, nil
}
