package interactor

import (
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type LikeInteractor struct {
	DB   repository.DBRepository
	Like repository.LikeRepository
}

func (i *LikeInteractor) Get(like *models.Likes) (*models.Likes, *services.ResultStatus) {
	db := i.DB.Connect()

	foundLike, err := i.Like.FindByID(db, like.ID)
	if err != nil {
		return &models.Likes{}, services.NewResultStatus(400, err)
	}

	return foundLike, services.NewResultStatus(200, nil)
}

// func (i *LikeInteractor) GetList() (*models.Likes, *services.ResultStatus) {

// }

func (i *LikeInteractor) Create(like *models.Likes) (*models.Likes, *services.ResultStatus) {
	db := i.DB.Connect()
	// 重複しないように確認する
	if _, err := i.Like.FindBySendUserIDAndReceiveUserID(db, like.SendUserID, like.ReceiveUserID); err == nil {
		return &models.Likes{}, services.NewResultStatus(400, errors.New("既にいいねしています"))
	}

	if like.SendUserID < 0 {
		return &models.Likes{}, services.NewResultStatus(400, errors.New("senderID not empty"))
	}
	if like.ReceiveUserID < 0 {
		return &models.Likes{}, services.NewResultStatus(400, errors.New("receiverID not empty"))
	}

	like.CreatedAt = services.SetNowDate()

	newLike, err := i.Like.Create(db, like)
	if err != nil {
		return &models.Likes{}, services.NewResultStatus(400, err)
	}

	return newLike, services.NewResultStatus(200, nil)
}
