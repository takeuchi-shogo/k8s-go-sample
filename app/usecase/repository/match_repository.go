package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type MatchRepository interface {
	// FindByUserID(db *gorm.DB, userID int) ([]*models.Matches, error)
	Create(db *gorm.DB, match *models.Matches) (*models.Matches, error)
}
