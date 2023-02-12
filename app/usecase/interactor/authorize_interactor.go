package interactor

import (
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/gateway"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type AuthorizeInteractor struct {
	DBRepository   repository.DBRepository
	Jwt            gateway.JwtGateway
	UserRepository repository.UserRepository
}

func (interactor *AuthorizeInteractor) Verify(token string) (*models.Users, *services.ResultStatus) {
	claims, err := interactor.Jwt.ParseToken(token)
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusUnauthorized, err)
	}

	// check token
	if isTokenExpire(int64(claims["exp"].(float64))) {
		return &models.Users{}, services.NewResultStatus(http.StatusUnauthorized, err)
	}

	db := interactor.DBRepository.Connect()

	foundUser, err := interactor.UserRepository.FindByID(db, claims["aud"].(int))
	if err != nil {
		return &models.Users{}, services.NewResultStatus(http.StatusUnauthorized, err)
	}
	return foundUser, services.NewResultStatus(http.StatusOK, nil)
}

func isTokenExpire(expireAt int64) bool {
	currentTime := time.Now().Unix()
	if expireAt < currentTime {
		return true
	}
	return false
}
