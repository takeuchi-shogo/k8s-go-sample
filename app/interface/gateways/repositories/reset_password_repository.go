package repositories

import (
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type ResetPasswordRepository struct{}

func (r *ResetPasswordRepository) FirstByToken(db *gorm.DB, token string) (*models.ResetPasswords, error) {
	resetPassword := &models.ResetPasswords{}
	if err := db.Where("token = ?", token).First(resetPassword).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &models.ResetPasswords{}, err
	}
	return resetPassword, nil
}

func (r *ResetPasswordRepository) FirstByResetKey(db *gorm.DB, resetKey string) (*models.ResetPasswords, error) {
	resetPassword := &models.ResetPasswords{}
	if err := db.Where("reset_key = ?", resetKey).First(resetPassword).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &models.ResetPasswords{}, err
	}
	return resetPassword, nil
}

func (r *ResetPasswordRepository) Create(db *gorm.DB, resetPassword *models.ResetPasswords) (*models.ResetPasswords, error) {
	if err := db.Create(resetPassword).Error; errors.Is(err, gorm.ErrRegistered) {
		return &models.ResetPasswords{}, err
	}
	return resetPassword, nil
}

func (r *ResetPasswordRepository) Save(db *gorm.DB, resetPassword *models.ResetPasswords) (*models.ResetPasswords, error) {
	if err := db.Save(resetPassword).Error; !errors.Is(err, nil) {
		return &models.ResetPasswords{}, err
	}
	return resetPassword, nil
}
