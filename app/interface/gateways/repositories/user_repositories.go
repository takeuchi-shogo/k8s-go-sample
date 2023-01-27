package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) FindByID(id int) (*models.Users, error) {
	return &models.Users{}, nil
}

func (repo *UserRepository) Create(u *models.Users) (*models.Users, error) {
	if err := repo.DB.Create(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user create: %w", err)
	}
	return u, nil
}
