package interactor

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/gateway"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
	"github.com/takeuchi-shogo/k8s-go-sample/utils"
)

type AuthorizeInteractor struct {
	AccountRepository repository.AccountRepository
	DBRepository      repository.DBRepository
	Jwt               gateway.JwtGateway
	UserRepository    repository.UserRepository
}

type AuthToken struct {
	JwtToken string
	User     *models.Users
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

func (interactor *AuthorizeInteractor) Create(user *models.Users) (auth AuthToken, resultStatus *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	foundUser, err := interactor.UserRepository.FirstByScreenName(db, user.ScreenName)
	if err != nil {
		return AuthToken{}, services.NewResultStatus(400, errors.New("user is not found"))
	}

	// foundAccount, err := interactor.AccountRepository.FindByID(foundUser.AccountID)
	// if err != nil {
	// 	return  AuthToken{}, services.NewResultStatus(400, []string{}, errors.New("user is not found"))
	// }

	// if utils.CheckPasswordHash(foundAccount.Password, foundUser.Password) {
	// 	return AuthToken{}, services.NewResultStatus(400, []string{}, errors.New("test auth error"))
	// }

	token := interactor.Jwt.CreateToken(foundUser.ID)

	return AuthToken{JwtToken: token, User: foundUser}, services.NewResultStatus(200, nil)
}

func (interactor *AuthorizeInteractor) Login(email, password string) (string, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	foundAccount, err := interactor.AccountRepository.FirstByEmail(db, email)
	if err != nil {
		return "", services.NewResultStatus(400, err)
	}

	if err = utils.CheckPassword(password, foundAccount.Password); err != nil {
		return "", services.NewResultStatus(400, err)
	}

	foundUser, err := interactor.UserRepository.FirstByAccountID(db, foundAccount.ID)
	if err != nil {
		return "", services.NewResultStatus(400, err)
	}

	token := interactor.Jwt.CreateToken(foundUser.ID)
	log.Println("5")

	return token, services.NewResultStatus(200, nil)
}

func isTokenExpire(expireAt int64) bool {
	currentTime := time.Now().Unix()
	if expireAt < currentTime {
		return true
	}
	return false
}
