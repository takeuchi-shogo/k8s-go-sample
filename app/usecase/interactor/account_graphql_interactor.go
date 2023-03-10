package interactor

import (
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
	"github.com/takeuchi-shogo/k8s-go-sample/utils"
)

type AccountGraphqlInteractor struct {
	Account          repository.AccountRepository
	DB               repository.DBRepository
	User             repository.UserRepository
	UserProfile      repository.UserProfileRepository
	UserSearchFilter repository.UserSearchFilterRepository
}

func (interactor *AccountGraphqlInteractor) Signup(account *models.Accounts, user *models.Users) (*models.Users, *services.ResultStatus) {

	db := interactor.DB.Begin()

	account.PhoneNumber = "000000"
	account.Password, _ = utils.GenerateFromPassword(account.Password)
	account.LoginStatus = "login"
	account.AccessLevel = 1
	currentTime := time.Now().Unix()
	account.CreatedAt = currentTime
	account.UpdatedAt = currentTime

	createdAccount, err := interactor.Account.Create(db, account)
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

	createdUser, err := interactor.User.Create(db, user)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	if _, err := interactor.UserProfile.Create(db, &models.UserProfiles{
		UserID: createdUser.ID,
	}); err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	if _, err := interactor.UserSearchFilter.Create(db, &models.UserSearchFilters{
		UserID: createdUser.ID,
	}); err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return createdUser, services.NewResultStatus(http.StatusOK, nil)
}
