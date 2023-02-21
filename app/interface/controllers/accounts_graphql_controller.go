package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type AccountsGraphqlController struct {
	Interactor interactor.AccountGraphqlInteractor
}

func NewAccountsGraphqlController(db repositories.DB) *AccountsGraphqlController {
	return &AccountsGraphqlController{
		Interactor: interactor.AccountGraphqlInteractor{
			Account: &repositories.AccountRepository{},
			DB:      &repositories.DBRepository{DB: db},
			User:    &repositories.UserRepository{},
		},
	}
}

func (controller *AccountsGraphqlController) Post(
	ctx context.Context,
	account *models.Accounts,
	user *models.Users) (*models.Users, error) {

	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	createdUser, res := controller.Interactor.Signup(account, user)
	if res.Error != nil {
		return createdUser, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	gc.Header("Authorization", "Bearer "+"jwt")
	return createdUser, nil
}
