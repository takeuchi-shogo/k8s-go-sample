package controllers

import (
	"context"
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type LikesGraphqlController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.LikeInteractor
}

func NewLikesGraphqlController(db repositories.DB, jwt gateways.Jwt) *LikesGraphqlController {
	return &LikesGraphqlController{
		Authorize: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
		Interactor: interactor.LikeInteractor{
			DB:   &repositories.DBRepository{DB: db},
			Like: &repositories.LikeRepository{},
		},
	}
}

func (controller *LikesGraphqlController) Get(ctx context.Context, id int) (*models.Likes, error) {
	like, res := controller.Interactor.Get(&models.Likes{
		ID: id,
	})
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return like, nil
}

func (controller *LikesGraphqlController) Post(ctx context.Context, like *models.Likes) (*models.Likes, error) {

	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token := gc.GetHeader("authorization")
	if token == "" {
		return nil, helpers.GraphQLErrorResponse(ctx, errors.New("ログインしてください"), 401)
	}

	userID, res := controller.Authorize.Verify(token)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	like.SendUserID = userID

	newLike, res := controller.Interactor.Create(like)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return newLike, nil
}
