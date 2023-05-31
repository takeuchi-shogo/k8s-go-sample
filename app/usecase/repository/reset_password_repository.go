package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type ResetPasswordRepository interface {
	FirstByToken(db *gorm.DB, token string) (*models.ResetPasswords, error)
	FirstByResetKey(db *gorm.DB, resetKey string) (*models.ResetPasswords, error)
	Create(db *gorm.DB, resetPassword *models.ResetPasswords) (*models.ResetPasswords, error)
	Save(db *gorm.DB, resetPassword *models.ResetPasswords) (*models.ResetPasswords, error)
}
