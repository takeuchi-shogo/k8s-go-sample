package repository

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"gorm.io/gorm"
)

type LikeRepository interface {
	FindByID(db *gorm.DB, id int) (*models.Likes, error)
	FindBySendUserIDAndReceiveUserID(db *gorm.DB, senderID, receiverID int) (*models.Likes, error)
	Create(db *gorm.DB, like *models.Likes) (*models.Likes, error)
}
