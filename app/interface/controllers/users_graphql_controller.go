package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/presenters"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type UsersGraphqlController struct {
	Interactor interactor.UserInteractor
}

func NewUsersGraphqlController(db repositories.DB) *UsersGraphqlController {
	return &UsersGraphqlController{
		Interactor: interactor.UserInteractor{
			AccountRepository:     &repositories.AccountRepository{},
			DBRepository:          &repositories.DBRepository{DB: db},
			UserRepository:        &repositories.UserRepository{},
			UserPresenter:         &presenters.UsersPresenter{},
			VerifyEmailRepository: &repositories.VerifyEmailRepository{},
		},
	}
}

func (controller *UsersGraphqlController) GetList(ctx context.Context, first int, after string, filter *types.UserFilter) (*types.UserConnection, error) {
	users, res := controller.Interactor.GetList(first, after, filter)
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
