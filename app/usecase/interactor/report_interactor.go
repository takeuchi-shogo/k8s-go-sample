package interactor

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type ReportInteractor struct {
	DBRepository     repository.DBRepository
	ReportRepository repository.ReportRepository
}

func (interactor *ReportInteractor) Create(r *models.Reports) (*models.Reports, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	report, err := interactor.ReportRepository.Create(db, r)
	if err != nil {
		return &models.Reports{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return report, services.NewResultStatus(http.StatusOK, nil)
}
