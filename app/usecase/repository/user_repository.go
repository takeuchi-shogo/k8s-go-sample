package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (*models.Users, error)
	Create(db *gorm.DB, u *models.Users) (*models.Users, error)
}
