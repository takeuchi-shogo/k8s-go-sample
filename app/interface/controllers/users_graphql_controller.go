package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
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

func (controller *UsersGraphqlController) GetList(ctx context.Context) ([]*models.Users, error) {
	users, res := controller.Interactor.GetList()
	if res.Error != nil {
		return users, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	// gc, err := helpers.GinContextFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// gc.Header("Authorization", "Bearer "+"testJwtToken")

	return users, nil
}

func (controller *UsersGraphqlController) Patch(ctx context.Context, user *models.Users) (*models.Users, error) {
	updatedUser, res := controller.Interactor.Save(user)
	if res.Error != nil {
		return user, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return updatedUser, nil
}
