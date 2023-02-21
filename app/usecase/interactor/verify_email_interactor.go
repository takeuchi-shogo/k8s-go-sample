package interactor

import (
	"errors"
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
	"github.com/takeuchi-shogo/k8s-go-sample/utils"
)

type VerifyEmailInteractor struct {
	DBRepository          repository.DBRepository
	VerifyEmailRepository repository.VerifyEmailRepository
}

func setVerifyEmail(email string) *models.VerifyEmails {
	v := &models.VerifyEmails{
		Email:    email,
		Token:    utils.RandomString(25),
		PINCode:  utils.RandomPinCode(),
		ExpireAt: time.Now().Add(1 * time.Hour).Unix(),
		CreateAt: time.Now().Unix(),
	}
	return v
}

func (interactor *VerifyEmailInteractor) GetByPINCode(pinCode string) (*models.VerifyEmails, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	verify, err := interactor.VerifyEmailRepository.FirstByCode(db, pinCode)
	if err != nil {
		return &models.VerifyEmails{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	if verify.ExpireAt < time.Now().Unix() {
		return &models.VerifyEmails{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return verify, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *VerifyEmailInteractor) Create(email string) (*models.VerifyEmails, *services.ResultStatus) {

	db := interactor.DBRepository.Connect()

	if _, err := interactor.VerifyEmailRepository.FirstByEmail(db, email); err == nil {
		return &models.VerifyEmails{}, services.NewResultStatus(http.StatusBadRequest, errors.New("既に使用されているメールアドレスです"))
	}

	verifyEmail := setVerifyEmail(email)

	createdVerifyEmail, err := interactor.VerifyEmailRepository.Create(db, verifyEmail)
	if err != nil {
		return &models.VerifyEmails{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	return createdVerifyEmail, services.NewResultStatus(http.StatusOK, nil)
}
