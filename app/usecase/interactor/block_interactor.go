package interactor

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type BlockInteractor struct {
	BlockRepository repository.BlockRepository
	DBRepository    repository.DBRepository
}

func (interactor *BlockInteractor) Create(b *models.Blocks) (*models.Blocks, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	block, err := interactor.BlockRepository.Create(db, b)
	if err != nil {
		return &models.Blocks{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return block, services.NewResultStatus(http.StatusOK, nil)
}
