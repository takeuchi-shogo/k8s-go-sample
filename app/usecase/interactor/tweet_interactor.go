package interactor

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type TweetInteractor struct {
	DB    repository.DBRepository
	Tweet repository.TweetRepository
}

func (i *TweetInteractor) Get(id int) (*models.Tweets, *services.ResultStatus) {

	db := i.DB.Connect()

	tweet, err := i.Tweet.TakeByID(db, id)
	if err != nil {
		return &models.Tweets{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	return tweet, services.NewResultStatus(http.StatusOK, nil)
}

func (i *TweetInteractor) GetList(userID int) ([]*models.Tweets, services.ResultStatus) {

	db := i.DB.Connect()

	tweets, err := i.Tweet.FindByUserID(db, userID)
	if err != nil {
		return []*models.Tweets{}, *services.NewResultStatus(http.StatusNotFound, err)
	}

	return tweets, *services.NewResultStatus(http.StatusOK, nil)
}

// func (i *TweetInteractor) Create(db *gorm.DB, )
