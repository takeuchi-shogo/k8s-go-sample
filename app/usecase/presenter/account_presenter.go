package presenter

import "github.com/takeuchi-shogo/k8s-go-sample/domain/models"

type AccountPresenter interface {
	ResponseAccount(a *models.Accounts) *models.Accounts
}
