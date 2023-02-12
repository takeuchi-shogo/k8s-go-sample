package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type blocksController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.BlockInteractor
}

type BlocksControllerProvider struct {
	DB  repositories.DB
	Jwt gateways.Jwt
}

func NewBlocksController(p BlocksControllerProvider) *blocksController {
	return &blocksController{
		Authorize: interactor.AuthorizeInteractor{
			Jwt: &gateways.JwtGateway{Jwt: p.Jwt},
		},
		Interactor: interactor.BlockInteractor{
			BlockRepository: &repositories.BlockRepository{},
		},
	}
}

func (controller *blocksController) Get(c helpers.Context) {

}

func (controller *blocksController) Post(c helpers.Context) {

}
