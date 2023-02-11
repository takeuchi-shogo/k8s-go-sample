package interactor

import "github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"

type ReportInteractor struct {
	ReportRepository repository.ReportRepository
}
