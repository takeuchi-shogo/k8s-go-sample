package controllers

import (
	"context"
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/presenters"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type UsersGraphqlController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.UserInteractor
}

func NewUsersGraphqlController(db repositories.DB, jwt gateways.Jwt) *UsersGraphqlController {
	return &UsersGraphqlController{
		Authorize: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
		Interactor: interactor.UserInteractor{
			AccountRepository:          &repositories.AccountRepository{},
			DBRepository:               &repositories.DBRepository{DB: db},
			UserRepository:             &repositories.UserRepository{},
			UserSearchFilterRepository: &repositories.UserSearchFilterRepository{},
			UserPresenter:              &presenters.UsersPresenter{},
			VerifyEmailRepository:      &repositories.VerifyEmailRepository{},
		},
	}
}

func (controller *UsersGraphqlController) GetList(ctx context.Context, first int, after string) (*types.UserConnection, error) {
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
		return &types.UserConnection{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	users, res := controller.Interactor.GetList(first, after, userID)
	if res.Error != nil {
		return &types.UserConnection{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	// gc, err := helpers.GinContextFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// gc.Header("Authorization", "Bearer "+"testJwtToken")

	return users, nil
}

func (controller *UsersGraphqlController) Get(ctx context.Context, id int) (*models.Users, error) {
	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		return user, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	// gc, err := helpers.GinContextFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// gc.Header("Authorization", "Bearer "+"testJwtToken")

	return user, nil
}

func (controller *UsersGraphqlController) Patch(ctx context.Context, user *models.Users) (*models.Users, error) {
	updatedUser, res := controller.Interactor.Save(user)
	if res.Error != nil {
		return user, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return updatedUser, nil
}
