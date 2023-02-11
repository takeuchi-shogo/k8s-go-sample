package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type ReportRepository struct{}

func (repo *ReportRepository) FindByID(db *gorm.DB, id int) (*models.Reports, error) {
	report := &models.Reports{}
	if err := db.First(report, id).Error; !errors.Is(err, nil) {
		return &models.Reports{}, fmt.Errorf("%w", err)
	}
	return report, nil
}

func (repo *ReportRepository) FindByUserID(db *gorm.DB, userID int) (*models.Reports, error) {
	report := &models.Reports{}
	if err := db.Where("user_id = ?", userID).Error; !errors.Is(err, nil) {
		return &models.Reports{}, fmt.Errorf("%w", err)
	}
	return report, nil
}

func (repo *ReportRepository) Create(db *gorm.DB, r *models.Reports) (*models.Reports, error) {
	report := &models.Reports{}
	report = r
	if err := db.Create(report).Error; !errors.Is(err, nil) {
		return &models.Reports{}, fmt.Errorf("%w", err)
	}
	return report, nil
}
