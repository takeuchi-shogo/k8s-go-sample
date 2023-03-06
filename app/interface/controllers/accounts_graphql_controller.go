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

type AccountsGraphqlController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.AccountInteractor
}

func NewAccountsGraphqlController(db repositories.DB, jwt gateways.Jwt) *AccountsGraphqlController {
	return &AccountsGraphqlController{
		Authorize: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
		Interactor: interactor.AccountInteractor{
			AccountRepository:          &repositories.AccountRepository{},
			DBRepository:               &repositories.DBRepository{DB: db},
			Jwt:                        &gateways.JwtGateway{Jwt: jwt},
			UserProfileRepository:      &repositories.UserProfileRepository{},
			UserRepository:             &repositories.UserRepository{},
			UserSearchFilterRepository: &repositories.UserSearchFilterRepository{},

			VerifyEmailRepository: &repositories.VerifyEmailRepository{},
		},
	}
}

func (controller *AccountsGraphqlController) Get(ctx context.Context, userID int) (*models.Accounts, error) {

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

	account, res := controller.Interactor.Get(userID)
	if res.Error != nil {
		return account, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return account, nil
}

func (controller *AccountsGraphqlController) Post(
	ctx context.Context,
	account *models.Accounts,
	user *models.Users) (*models.Users, error) {

	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// Account自体は返さない
	createdUser, jwtToken, res := controller.Interactor.Signup(account, user)
	if res.Error != nil {
		return createdUser, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	gc.Header("Authorization", "Bearer "+jwtToken)
	return createdUser, nil
}

func (controller *AccountsGraphqlController) Patch(
	ctx context.Context,
	account *models.Accounts,
) (*models.Accounts, error) {
	updatedAccount, res := controller.Interactor.Save(account)
	if res.Error != nil {
		return updatedAccount, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return updatedAccount, nil
}
