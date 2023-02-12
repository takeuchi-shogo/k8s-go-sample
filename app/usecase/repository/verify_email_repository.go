package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type VerifyEmailRepository interface {
	FirstByCode(db *gorm.DB, code string) (*models.VerifyEmails, error)
	FirstByToken(db *gorm.DB, token string) (*models.VerifyEmails, error)
	FirstByEmail(db *gorm.DB, email string) (*models.VerifyEmails, error)
	Create(db *gorm.DB, v *models.VerifyEmails) (*models.VerifyEmails, error)
}
