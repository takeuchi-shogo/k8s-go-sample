package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type userProfilesController struct {
	AuthorizeInteractor interactor.AuthorizeInteractor
	Interactor          interactor.UserProfileInteractor
}

type UserProfileInteractor interface {
	Get()
	Post()
}

type UserProfilesControllerProvider struct {
	DB  repositories.DB
	Jwt gateways.Jwt
}

func NewUserProfilesController(p UserProfilesControllerProvider) *userProfilesController {
	return &userProfilesController{
		AuthorizeInteractor: interactor.AuthorizeInteractor{
			Jwt: &gateways.JwtGateway{Jwt: p.Jwt},
		},
		Interactor: interactor.UserProfileInteractor{
			DBRepository:          &repositories.DBRepository{DB: p.DB},
			UserProfileRepository: &repositories.UserProfileRepository{},
		},
	}
}

func (controller *userProfilesController) Get(c helpers.Context) {

}

func (controller *userProfilesController) Post(c helpers.Context) {

}

func (controller *userProfilesController) Patch(c helpers.Context) {

}
