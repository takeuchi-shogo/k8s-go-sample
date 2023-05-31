package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type TweetRepository interface {
	TakeByID(db *gorm.DB, id int) (*models.Tweets, error)
	FindByUserID(db *gorm.DB, userID int) ([]*models.Tweets, error)
	LastByUserID(db *gorm.DB, userID int) (*models.Tweets, error)
	Create(db *gorm.DB, tweet *models.Tweets) (*models.Tweets, error)
	Save(db *gorm.DB, tweet *models.Tweets) (*models.Tweets, error)
	Delete(db *gorm.DB, id int) error
}
