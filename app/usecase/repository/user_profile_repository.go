package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	FindByID(db *gorm.DB, id int) (*models.UserProfiles, error)
	FindByUserID(db *gorm.DB, userID int) (*models.UserProfiles, error)
	Create(db *gorm.DB, p *models.UserProfiles) (*models.UserProfiles, error)
	Save(db *gorm.DB, p *models.UserProfiles) (*models.UserProfiles, error)
}
