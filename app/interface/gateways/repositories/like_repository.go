package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type LikeRepository struct{}

func (repo *LikeRepository) FindByID(db *gorm.DB, id int) (*models.Likes, error) {
	like := &models.Likes{}
	if err := db.First(like, id).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed get like by id: %w", err)
	}
	return like, nil
}

func (repo *LikeRepository) FindBySendUserIDAndReceiveUserID(db *gorm.DB, senderID, receiverID int) (*models.Likes, error) {
	like := &models.Likes{}
	if err := db.Where("send_user_id = ? and receive_user_id = ?", senderID, receiverID).First(like).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed get like by send_user_id and receive_user_id: %w", err)
	}
	return like, nil
}

func (repo *LikeRepository) Create(db *gorm.DB, like *models.Likes) (*models.Likes, error) {
	if err := db.Create(like).Error; !errors.Is(err, nil) {
		return nil, fmt.Errorf("failed create like: %w", err)
	}
	return like, nil
}
