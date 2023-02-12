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
		Token:    utils.RandomString(15),
		PINCode:  services.CreatePINCode(8),
		ExpireAt: time.Now().Add(1 * time.Hour).Unix(),
	}
	return v
}

func (interactor *VerifyEmailInteractor) GetByPINCode(pinCode string) *services.ResultStatus {
	db := interactor.DBRepository.Connect()

	verify, err := interactor.VerifyEmailRepository.FirstByCode(db, pinCode)
	if err != nil {
		return services.NewResultStatus(http.StatusBadRequest, err)
	}
	if verify.ExpireAt < time.Now().Unix() {
		return services.NewResultStatus(http.StatusBadRequest, err)
	}
	return services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *VerifyEmailInteractor) Create(email string) *services.ResultStatus {

	db := interactor.DBRepository.Connect()

	if _, err := interactor.VerifyEmailRepository.FirstByEmail(db, email); err == nil {
		return services.NewResultStatus(http.StatusBadRequest, errors.New("既に使用されているメールアドレスです"))
	}

	verifyEmail := setVerifyEmail(email)

	_, err := interactor.VerifyEmailRepository.Create(db, verifyEmail)
	if err != nil {
		return services.NewResultStatus(http.StatusBadRequest, err)
	}

	return services.NewResultStatus(http.StatusOK, nil)
}
