package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserSearchFilterRepository interface {
	FindByID(db *gorm.DB, id int) (*models.UserSearchFilters, error)
	FirstByUserID(db *gorm.DB, userID int) (*models.UserSearchFilters, error)
	Create(db *gorm.DB, filter *models.UserSearchFilters) (*models.UserSearchFilters, error)
	Save(db *gorm.DB, filter *models.UserSearchFilters) (*models.UserSearchFilters, error)
}
