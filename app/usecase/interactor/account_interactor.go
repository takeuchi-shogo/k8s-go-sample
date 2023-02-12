package interactor

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type AccountInteractor struct {
	AccountRepository     repository.AccountRepository
	DBRepository          repository.DBRepository
	VerifyEmailRepository repository.VerifyEmailRepository

	AccountPresenter presenter.AccountPresenter
}

func (interactor *AccountInteractor) Get(id int) (*models.Accounts, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	account, err := interactor.AccountRepository.FindByID(db, id)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return account, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *AccountInteractor) Create(a *models.Accounts, token string) (*models.Accounts, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	verify, err := interactor.VerifyEmailRepository.FirstByToken(db, token)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	a.Email = verify.Email

	account, err := interactor.AccountRepository.Create(db, a)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return account, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *AccountInteractor) Save(a *models.Accounts) (*models.Accounts, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	account, err := interactor.AccountRepository.Save(db, a)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return account, services.NewResultStatus(http.StatusOK, nil)
}
