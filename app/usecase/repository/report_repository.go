package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type ReportRepository interface {
	FindByID(db *gorm.DB, id int) (*models.Reports, error)
	FindByUserID(db *gorm.DB, userID int) (*models.Reports, error)
	Create(db *gorm.DB, report *models.Reports) (*models.Reports, error)
}
