package controllers

import (
	"context"
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/presenters"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type MeGraphqlController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.MeInteractor
}

func NewMeGraphqlController(db repositories.DB, jwt gateways.Jwt) *MeGraphqlController {
	return &MeGraphqlController{
		Authorize: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
		Interactor: interactor.MeInteractor{
			DB:                   &repositories.DBRepository{DB: db},
			Like:                 &repositories.LikeRepository{},
			User:                 &repositories.UserRepository{},
			UserProfile:          &repositories.UserProfileRepository{},
			UserProfilePresenter: &presenters.UserPriflesPresenter{},
		},
	}
}

func (controller *MeGraphqlController) Get(ctx context.Context) (*models.ResponseUsers, error) {

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

	me, res := controller.Interactor.Get(userID)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return me, nil
}
