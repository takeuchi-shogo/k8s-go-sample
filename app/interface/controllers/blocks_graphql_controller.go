package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type BlocksGraphqlController struct {
	Interactor interactor.BlockInteractor
}

func NewBlocksGraphqlController(db repositories.DB) *BlocksGraphqlController {
	return &BlocksGraphqlController{
		Interactor: interactor.BlockInteractor{},
	}
}

func (controller *BlocksGraphqlController) Post(ctx context.Context, block *models.Blocks) (*models.Blocks, error) {
	newBlock, res := controller.Interactor.Create(block)
	if res.Error != nil {
		return nil, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return newBlock, nil
}
