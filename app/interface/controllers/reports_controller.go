package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type reportsController struct {
	AuthorizeInteractor interactor.AuthorizeInteractor
	Interactor          interactor.ReportInteractor
}

type ReportsControllerProvider struct {
	DB  repositories.DB
	Jwt gateways.Jwt
}

func NewReportsController(p ReportsControllerProvider) *reportsController {
	return &reportsController{
		AuthorizeInteractor: interactor.AuthorizeInteractor{
			Jwt: &gateways.JwtGateway{Jwt: p.Jwt},
		},
		Interactor: interactor.ReportInteractor{},
	}
}

func (controller *reportsController) Get(c helpers.Context) {

}

func (controller *reportsController) Post(c helpers.Context) {

}
