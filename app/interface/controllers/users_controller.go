package controllers

import (
	"log"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/presenters"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type usersController struct {
	Interactor interactor.UserInteractor
}

type UserController interface {
	Get()
	Post()
}

func NewUsersController(db repositories.DB) *usersController {
	return &usersController{
		Interactor: interactor.UserInteractor{
			UserRepository: &repositories.UserRepository{},
			UserPresenter:  &presenters.UsersPresenter{},
		},
	}
}

func (controller *usersController) Get(c helpers.Context) {
	log.Println("get user")
}

func (controller *usersController) Post(c helpers.Context) {
	u := &models.Users{}
	user, res := controller.Interactor.Create(u)
	if res.Error != nil {
		c.JSON(res.Code, helpers.NewResponseError(res.Code, res.Error, res.Error.Error()))
	}
	c.JSON(res.Code, helpers.NewResponseSuccess(user, "success"))
}
