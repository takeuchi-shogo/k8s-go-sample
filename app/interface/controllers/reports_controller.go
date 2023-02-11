package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type reportsController struct {
	Interactor interactor.ReportInteractor
}

func NewReportsController(db repositories.DB) *reportsController {
	return &reportsController{
		Interactor: interactor.ReportInteractor{},
	}
}

func (controller *reportsController) Get(c helpers.Context) {

}

func (controller *reportsController) Post(c helpers.Context) {

}
