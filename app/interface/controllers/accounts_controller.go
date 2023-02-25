package controllers

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type accountsController struct {
	AuthorizeInteractor interactor.AuthorizeInteractor
	Interactor          interactor.AccountInteractor
}

type AccountsController interface {
	Get()
	Post()
}

type AccountsControllerProvider struct {
	DB  repositories.DB
	Jwt gateways.Jwt
}

func NewAccountsController(p AccountsControllerProvider) *accountsController {
	return &accountsController{
		AuthorizeInteractor: interactor.AuthorizeInteractor{
			Jwt: &gateways.JwtGateway{Jwt: p.Jwt},
		},
		Interactor: interactor.AccountInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: p.DB},
			// AccountPresenter:  &presenters.AccountsPresenter{},
		},
	}
}

func (controller *accountsController) Get(c helpers.Context) {
	user, res := controller.AuthorizeInteractor.Verify(c.GetHeader("Authorize"))
	if res.Error != nil {
		return
	}
	c.JSON(res.Code, helpers.NewResponseSuccess("success", user))
}

// Create Account
func (controller *accountsController) Post(c helpers.Context) {
	token := c.PostForm("token")

	a := &models.Accounts{}

	if err := c.BindJSON(a); err != nil {
		code := http.StatusBadRequest
		c.JSON(code, helpers.NewResponseError(code, err, err.Error()))
		return
	}

	account, res := controller.Interactor.Create(a, token)
	if res.Error != nil {
		c.JSON(res.Code, helpers.NewResponseError(res.Code, res.Error, res.Error.Error()))
		return
	}
	c.JSON(res.Code, helpers.NewResponseSuccess("success", account))
}
