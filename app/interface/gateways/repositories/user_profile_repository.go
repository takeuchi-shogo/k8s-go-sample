package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type UserProfileRepository struct{}

func (repo *UserProfileRepository) FindByID(db *gorm.DB, id int) (*models.UserProfiles, error) {
	profile := &models.UserProfiles{}
	if err := db.First(profile, id).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed find profile by id : %w", err)
	}
	return profile, nil
}

func (repo *UserProfileRepository) FindByUserID(db *gorm.DB, userID int) (*models.UserProfiles, error) {
	profile := &models.UserProfiles{}
	if err := db.Where("user_id = ?", userID).First(profile).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed find profile by user id : %w", err)
	}
	return profile, nil
}

func (repo *UserProfileRepository) Create(db *gorm.DB, p *models.UserProfiles) (*models.UserProfiles, error) {
	profile := &models.UserProfiles{}
	profile = p
	if err := db.Create(profile).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed profile create: %w", err)
	}
	return profile, nil
}

func (repo *UserProfileRepository) Save(db *gorm.DB, p *models.UserProfiles) (*models.UserProfiles, error) {
	profile := &models.UserProfiles{}
	profile = p
	if err := db.Save(profile).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed profile save: %w", err)
	}
	return profile, nil
}
