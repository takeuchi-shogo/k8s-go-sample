package interactor

import (
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
	"github.com/takeuchi-shogo/k8s-go-sample/utils"
)

type AccountInteractor struct {
	AccountRepository     repository.AccountRepository
	DBRepository          repository.DBRepository
	UserRepository        repository.UserRepository
	UserProfileRepository repository.UserProfileRepository
	VerifyEmailRepository repository.VerifyEmailRepository

	// AccountPresenter presenter.AccountPresenter
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

func (interactor *AccountInteractor) Signup(account *models.Accounts, user *models.Users) (*models.Users, *services.ResultStatus) {
	db := interactor.DBRepository.Begin()

	account.PhoneNumber = "000000"
	account.Password, _ = utils.GenerateFromPassword(account.Password)
	account.LoginStatus = "login"
	account.AccessLevel = 1
	currentTime := time.Now().Unix()
	account.CreatedAt = currentTime
	account.UpdatedAt = currentTime

	createdAccount, err := interactor.AccountRepository.Create(db, account)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	user.AccountID = createdAccount.ID
	user.UUID = utils.NewRandom()
	user.ScreenName = utils.RandomScreenName()
	user.IsAuthorizeEmail = false
	user.IsVerifiedAge = false
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	createdUser, err := interactor.UserRepository.Create(db, user)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	if _, err := interactor.UserProfileRepository.Create(db, &models.UserProfiles{
		UserID:    createdUser.ID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}); err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return createdUser, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *AccountInteractor) Save(a *models.Accounts) (*models.Accounts, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	account, err := interactor.AccountRepository.Save(db, a)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return account, services.NewResultStatus(http.StatusOK, nil)
}
