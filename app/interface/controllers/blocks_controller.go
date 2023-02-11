package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type blocksController struct {
	Interactor interactor.BlockInteractor
}

func NewBlocksController(db repositories.DB) *blocksController {
	return &blocksController{
		Interactor: interactor.BlockInteractor{
			BlockRepository: &repositories.BlockRepository{},
		},
	}
}

func (controller *blocksController) Get(c helpers.Context) {

}

func (controller *blocksController) Post(c helpers.Context) {

}
