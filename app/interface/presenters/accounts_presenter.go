package presenters

import "github.com/takeuchi-shogo/k8s-go-sample/domain/models"

type AccountsPresenter struct{}

func (p *AccountsPresenter) ResponseAccount(a *models.Accounts) *models.Accounts {
	return a
}
