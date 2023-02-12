package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (*models.Users, error) {
	return &models.Users{}, nil
}

func (repo *UserRepository) Create(db *gorm.DB, u *models.Users) (*models.Users, error) {
	if err := db.Create(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user create: %w", err)
	}
	return u, nil
}

func (repo *UserRepository) Save(db *gorm.DB, u *models.Users) (*models.Users, error) {
	if err := db.Save(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user save: %w", err)
	}
	return u, nil
}
