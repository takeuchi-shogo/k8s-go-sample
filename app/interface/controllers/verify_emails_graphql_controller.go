package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type VerifyEmailsGraphqlController struct {
	Interactor interactor.VerifyEmailInteractor
}

func NewVerifyEmailsGraphqlController(db repositories.DB) *VerifyEmailsGraphqlController {
	return &VerifyEmailsGraphqlController{
		Interactor: interactor.VerifyEmailInteractor{
			DBRepository:          &repositories.DBRepository{DB: db},
			VerifyEmailRepository: &repositories.VerifyEmailRepository{},
		},
	}
}

func (controller *VerifyEmailsGraphqlController) Get(ctx context.Context, code string) (*models.VerifyEmails, error) {
	verifyEmail, res := controller.Interactor.GetByPINCode(code)
	if res.Error != nil {
		return verifyEmail, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return verifyEmail, nil
}

func (controller *VerifyEmailsGraphqlController) Post(ctx context.Context, email string) (*models.VerifyEmails, error) {
	createdVerifyEmail, res := controller.Interactor.Create(email)
	if res.Error != nil {
		return createdVerifyEmail, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return createdVerifyEmail, nil
}
