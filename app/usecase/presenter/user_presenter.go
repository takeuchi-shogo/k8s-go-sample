package presenter

import "github.com/takeuchi-shogo/k8s-go-sample/domain/models"

type UserPresenter interface {
	ResponseUser(u *models.ResponseUsers) *models.ResponseUsers
}
