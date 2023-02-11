package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/presenters"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type accountsController struct {
	Interactor interactor.AccountInteractor
}

type AccountsController interface {
	Get()
	Post()
}

func NewAccountsController(db repositories.DB) *accountsController {
	return &accountsController{
		Interactor: interactor.AccountInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			AccountPresenter:  &presenters.AccountsPresenter{},
		},
	}
}

func (controller *accountsController) Get(c helpers.Context) {

}

func (controller *accountsController) Post(c helpers.Context) {

}
