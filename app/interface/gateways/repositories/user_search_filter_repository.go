package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserSearchFilterRepository struct{}

func (repo *UserSearchFilterRepository) FindByID(db *gorm.DB, id int) (*models.UserSearchFilters, error) {
	filter := &models.UserSearchFilters{}
	if err := db.First(filter, id).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed find filter by id: %w", err)
	}
	return filter, nil
}

func (repo *UserSearchFilterRepository) FirstByUserID(db *gorm.DB, userID int) (*models.UserSearchFilters, error) {
	filter := &models.UserSearchFilters{}
	if err := db.Where("user_id = ?", userID).First(filter).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed find filter by user id: %w", err)
	}
	return filter, nil
}

func (repo *UserSearchFilterRepository) Create(db *gorm.DB, filter *models.UserSearchFilters) (*models.UserSearchFilters, error) {
	// filter := &models.UserSearchFilters{}
	if err := db.Create(filter).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed create filter: %w", err)
	}
	return filter, nil
}

func (repo *UserSearchFilterRepository) Save(db *gorm.DB, filter *models.UserSearchFilters) (*models.UserSearchFilters, error) {
	if err := db.Save(filter).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed save filter: %w", err)
	}
	return filter, nil
}
