package interactor

import (
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type UserInteractor struct {
	AccountRepository     repository.AccountRepository
	DBRepository          repository.DBRepository
	UserRepository        repository.UserRepository
	UserPresenter         presenter.UserPresenter
	VerifyEmailRepository repository.VerifyEmailRepository
}

func (interactor *UserInteractor) GetList() ([]*models.Users, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	users, err := interactor.UserRepository.Find(db)
	if err != nil {
		return []*models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}
	return users, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *UserInteractor) Get(id int) (*models.Users, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	user, err := interactor.UserRepository.FindByID(db, id)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}
	return interactor.UserPresenter.ResponseUser(user), services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *UserInteractor) Create(u *models.Users, accountID int, code string) (*models.Users, *services.ResultStatus) {
	// repository process
	db := interactor.DBRepository.Connect()

	account, err := interactor.AccountRepository.FindByID(db, accountID)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	u.AccountID = account.ID

	if _, err := interactor.VerifyEmailRepository.FirstByCode(db, code); err == nil {
		u.IsAuthorizeEmail = true
	}

	user, err := interactor.UserRepository.Create(db, u)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	return user, services.NewResultStatus(http.StatusOK, nil)
}

func setValue(user, foundUser *models.Users) *models.Users {
	foundUser.DisplayName = user.DisplayName
	foundUser.ScreenName = user.ScreenName
	foundUser.Gender = user.Gender
	foundUser.Location = user.Location
	foundUser.UpdatedAt = time.Now().Unix()

	return foundUser
}

func (interactor *UserInteractor) Save(u *models.Users) (*models.Users, *services.ResultStatus) {

	db := interactor.DBRepository.Connect()

	foundUser, err := interactor.UserRepository.FindByID(db, u.ID)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	updatedUser := setValue(u, foundUser)

	user, err := interactor.UserRepository.Save(db, updatedUser)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	return user, services.NewResultStatus(http.StatusOK, nil)
}
