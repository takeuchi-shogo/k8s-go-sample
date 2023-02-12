package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type verifyEmailsController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.VerifyEmailInteractor
}

type VerifyEmailsControllerProvider struct {
	DB  repositories.DB
	Jwt gateways.Jwt
}

func NewVerifyEmailsController(p VerifyEmailsControllerProvider) *verifyEmailsController {
	return &verifyEmailsController{
		Authorize:  interactor.AuthorizeInteractor{},
		Interactor: interactor.VerifyEmailInteractor{},
	}
}

func (controller *verifyEmailsController) GetByPINCode(c helpers.Context) {
	pinCode := c.Query("pin_code")

	res := controller.Interactor.GetByPINCode(pinCode)
	if res.Error != nil {
		return
	}
	c.JSON(res.Code, helpers.NewResponseSuccess("success", nil))

}

func (controller *verifyEmailsController) Post(c helpers.Context) {
	email := c.PostForm("email")

	res := controller.Interactor.Create(email)
	if res.Error != nil {
		return
	}
	c.JSON(res.Code, helpers.NewResponseSuccess("success", nil))
}
