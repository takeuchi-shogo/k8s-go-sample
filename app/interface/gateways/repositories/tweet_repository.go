package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type TweetRepository struct{}

func (r *TweetRepository) TakeByID(db *gorm.DB, id int) (*models.Tweets, error) {
	t := &models.Tweets{}
	if err := db.Where("id = ?", id).Take(t).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &models.Tweets{}, fmt.Errorf("failed get tweet: %w\n", err)
	}
	return t, nil
}

func (r *TweetRepository) FindByUserID(db *gorm.DB, userID int) ([]*models.Tweets, error) {
	t := []*models.Tweets{}
	if err := db.Where("usre_id = ?", userID).Find(&t).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return []*models.Tweets{}, fmt.Errorf("failed get tweet list: %w\n", err)
	}
	return t, nil
}

func (r *TweetRepository) LastByUserID(db *gorm.DB, userID int) (*models.Tweets, error) {
	t := &models.Tweets{}
	if err := db.Where("usre_id = ?", userID).Last(&t).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &models.Tweets{}, fmt.Errorf("failed get tweet: %w\n", err)
	}
	return t, nil
}

func (r *TweetRepository) Create(db *gorm.DB, tweet *models.Tweets) (*models.Tweets, error) {
	if err := db.Create(tweet).Error; errors.Is(err, gorm.ErrRegistered) {
		return &models.Tweets{}, fmt.Errorf("failed create tweet: %w\n", err)
	}
	return tweet, nil
}

func (r *TweetRepository) Save(db *gorm.DB, tweet *models.Tweets) (*models.Tweets, error) {
	if err := db.Save(tweet).Error; !errors.Is(err, nil) {
		return &models.Tweets{}, fmt.Errorf("failed save tweet: %w\n", err)
	}
	return tweet, nil
}

func (r *TweetRepository) Delete(db *gorm.DB, id int) error {
	tweet, err := r.TakeByID(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed get tweet: %w\n", err)
	}

	deletedAt := time.Now().Unix()
	tweet.DeletedAt = &deletedAt

	if err := db.Save(tweet).Error; !errors.Is(err, nil) {
		return fmt.Errorf("failed delete tweet: %w\n", err)
	}

	return nil
}
