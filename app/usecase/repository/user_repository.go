package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(db *gorm.DB) ([]*models.Users, error)
	FindByID(db *gorm.DB, id int) (*models.Users, error)
	FirstByAccountID(db *gorm.DB, accountID int) (*models.Users, error)
	FirstByScreenName(db *gorm.DB, screenName string) (*models.Users, error)
	Create(db *gorm.DB, u *models.Users) (*models.Users, error)
	Save(db *gorm.DB, u *models.Users) (*models.Users, error)
}
