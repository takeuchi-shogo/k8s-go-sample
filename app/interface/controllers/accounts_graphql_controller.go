package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
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
	account *models.Accounts,
	user *models.Users) (*models.Users, error) {
	createdUser, res := controller.Interactor.Signup(account, user)
	return createdUser, res
}
