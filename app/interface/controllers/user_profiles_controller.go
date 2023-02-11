package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type userProfilesController struct {
	Interactor interactor.UserProfileInteractor
}

type UserProfileInteractor interface {
	Get()
	Post()
}

func NewUserProfilesController(db repositories.DB) *userProfilesController {
	return &userProfilesController{
		Interactor: interactor.UserProfileInteractor{},
	}
}

func (controller *userProfilesController) Get(c helpers.Context) {

}

func (controller *userProfilesController) Post(c helpers.Context) {

}
