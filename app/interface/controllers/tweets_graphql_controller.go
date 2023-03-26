package controllers

import (
	"context"
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type TweetsGraphqlController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.TweetInteractor
}

func NewTweetsGraphqlController(db repositories.DB, jwt gateways.Jwt) *TweetsGraphqlController {
	return &TweetsGraphqlController{
		Authorize: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
		Interactor: interactor.TweetInteractor{
			DB:    &repositories.DBRepository{DB: db},
			Tweet: &repositories.TweetRepository{},
		},
	}
}

func (c *TweetsGraphqlController) GetList(ctx context.Context, userID int) (*types.TweetConnection, error) {
	tweetConnection, res := c.Interactor.GetList(userID)
	if res.Error != nil {
		return &types.TweetConnection{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return tweetConnection, nil
}

func (c *TweetsGraphqlController) Create(ctx context.Context, text string) (*models.Tweets, error) {
	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token := gc.GetHeader("authorization")
	if token == "" {
		return nil, helpers.GraphQLErrorResponse(ctx, errors.New("ログインしてください"), 401)
	}

	userID, res := c.Authorize.Verify(token)
	if res.Error != nil {
		return &models.Tweets{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	tweet, res := c.Interactor.Create(&models.Tweets{
		UserID: userID,
		Text:   text,
	})
	if res.Error != nil {
		return &models.Tweets{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return tweet, nil
}
