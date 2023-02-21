package repositories

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type VerifyEmailRepository struct{}

func (repo *VerifyEmailRepository) FirstByCode(db *gorm.DB, code string) (*models.VerifyEmails, error) {
	return &models.VerifyEmails{}, nil
}
func (repo *VerifyEmailRepository) FirstByToken(db *gorm.DB, token string) (*models.VerifyEmails, error) {
	return &models.VerifyEmails{}, nil
}
func (repo *VerifyEmailRepository) FirstByEmail(db *gorm.DB, email string) (*models.VerifyEmails, error) {
	return &models.VerifyEmails{}, nil
}
func (repo *VerifyEmailRepository) Create(db *gorm.DB, v *models.VerifyEmails) (*models.VerifyEmails, error) {
	return &models.VerifyEmails{}, nil
}
