package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type ResetPasswordsGraphqlController struct {
	Interactor interactor.ResetPasswordInteractor
}

func NewResetPasswordsGraphqlController(db repositories.DB) *ResetPasswordsGraphqlController {
	return &ResetPasswordsGraphqlController{
		Interactor: interactor.ResetPasswordInteractor{
			Account:       &repositories.AccountRepository{},
			DB:            &repositories.DBRepository{DB: db},
			ResetPassword: &repositories.ResetPasswordRepository{},
			User:          &repositories.UserRepository{},
		},
	}
}

func (c *ResetPasswordsGraphqlController) Get(ctx context.Context, token string) (*models.ResetPasswords, error) {
	r, res := c.Interactor.Get(token)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return r, nil
}

func (c *ResetPasswordsGraphqlController) Create(ctx context.Context, email string) (*models.ResetPasswords, error) {
	r, res := c.Interactor.Create(email)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return r, nil
}

func (c *ResetPasswordsGraphqlController) Save(ctx context.Context, resetKey, password string) (*models.Users, error) {
	r, res := c.Interactor.Save(resetKey, password)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return r, nil
}
