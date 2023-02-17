package interactor

import (
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
)

type AccountGraphqlInteractor struct {
	Account repository.AccountRepository
	DB      repository.DBRepository
	User    repository.UserRepository
}

func (interactor *AccountGraphqlInteractor) Signup(account *models.Accounts, user *models.Users) (*models.Users, error) {

	db := interactor.DB.Begin()

	account.PhoneNumber = "000000"
	account.LoginStatus = "login"
	account.AccessLevel = 1
	currentTime := time.Now().Unix()
	account.CreatedAt = currentTime
	account.UpdatedAt = currentTime

	createdAccount, err := interactor.Account.Create(db, account)
	if err != nil {
		db.Rollback()
		return &models.Users{}, err
	}

	user.AccountID = createdAccount.ID
	user.IsAuthorizeEmail = false
	user.IsVerifiedAge = false
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	createdUser, err := interactor.User.Create(db, user)
	if err != nil {
		db.Rollback()
		return &models.Users{}, err
	}

	db.Commit()
	return createdUser, nil
}
