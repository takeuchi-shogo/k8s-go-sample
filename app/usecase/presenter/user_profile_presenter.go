package presenter

import "github.com/takeuchi-shogo/k8s-go-sample/domain/models"

type UserProfilePresenter interface {
	ResponseUserProfile(*models.UserProfiles) *models.UserProfiles
}
