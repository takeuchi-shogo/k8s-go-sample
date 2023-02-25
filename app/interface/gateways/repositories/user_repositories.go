package repositories

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) Find(db *gorm.DB) ([]*models.Users, error) {
	users := []*models.Users{}
	if err := db.Find(&users).Error; !errors.Is(err, nil) {
		return []*models.Users{}, fmt.Errorf("failed user get: %w", err)
	}
	return users, nil
}

func (repo *UserRepository) FindByUserFilter(db *gorm.DB, filter *types.UserFilter, limit int, after string) ([]*models.Users, error) {
	users := []*models.Users{}

	query := db.Order("created_at desc")

	if filter != nil {
		if *filter.Gender != "" {
			query = query.Where("gender = ?", *filter.Gender)
		}
		if *filter.Location != "" {
			query = query.Where("location = ?", *filter.Location)
		}
	}

	if after != "" {
		id, _ := strconv.Atoi(after)
		query = query.Where("id < ?", id)
	}

	query = query.Limit(limit).Find(&users)

	return users, query.Error
}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (*models.Users, error) {
	u := &models.Users{}
	if err := db.First(u, id).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user create: %w", err)
	}
	return u, nil
}

func (repo *UserRepository) FirstByAccountID(db *gorm.DB, accountID int) (*models.Users, error) {
	u := &models.Users{}
	if err := db.Where("account_id = ?", accountID).First(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user create: %w", err)
	}
	return u, nil
}

func (repo *UserRepository) FirstByScreenName(db *gorm.DB, screenName string) (*models.Users, error) {
	u := &models.Users{}
	if err := db.Where("screen_name = ?", screenName).First(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user create: %w", err)
	}
	return u, nil
}

func (repo *UserRepository) Create(db *gorm.DB, u *models.Users) (*models.Users, error) {
	if err := db.Create(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user create: %w", err)
	}
	return u, nil
}

func (repo *UserRepository) Save(db *gorm.DB, u *models.Users) (*models.Users, error) {
	if err := db.Save(u).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed user save: %w", err)
	}
	return u, nil
}
