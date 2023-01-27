package interactor

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

func (interactor *UserInteractor) Get(id int) (*models.Users, *services.ResultStatus) {
	user, err := interactor.UserRepository.FindByID(id)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}
	return user, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *UserInteractor) Create(u *models.Users) (*models.Users, *services.ResultStatus) {
	// repository prosess
	return u, services.NewResultStatus(http.StatusOK, nil)
}
