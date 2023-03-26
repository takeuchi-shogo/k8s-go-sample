package interactor

import (
	"net/http"
	"strconv"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
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

func (i *TweetInteractor) GetList(userID int) (*types.TweetConnection, services.ResultStatus) {

	db := i.DB.Connect()

	tweets, err := i.Tweet.FindByUserID(db, userID)
	if err != nil {
		return &types.TweetConnection{}, *services.NewResultStatus(http.StatusNotFound, err)
	}

	edges := []*types.TweetsEdge{}

	for _, tweet := range tweets {
		tweetEdges := &types.TweetsEdge{
			Cursor: strconv.Itoa(tweet.ID),
			Node:   tweet,
		}
		edges = append(edges, tweetEdges)
	}

	pageInfo := &types.PageInfo{}

	return &types.TweetConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}, *services.NewResultStatus(http.StatusOK, nil)
}

func (i *TweetInteractor) Create(tweet *models.Tweets) (*models.Tweets, *services.ResultStatus) {

	db := i.DB.Connect()

	tweet.CreatedAt = time.Now().Unix()
	tweet.UpdatedAt = time.Now().Unix()

	newTweet, err := i.Tweet.Create(db, tweet)
	if err != nil {
		return &models.Tweets{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return newTweet, services.NewResultStatus(http.StatusOK, nil)
}
