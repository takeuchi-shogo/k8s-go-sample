package interactor

import (
	"fmt"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
	"github.com/takeuchi-shogo/k8s-go-sample/utils"
)

type ResetPasswordInteractor struct {
	Account       repository.AccountRepository
	DB            repository.DBRepository
	ResetPassword repository.ResetPasswordRepository
	User          repository.UserRepository
}

func (i *ResetPasswordInteractor) Get(token string) (*models.ResetPasswords, *services.ResultStatus) {

	db := i.DB.Connect()

	foundResetPassword, err := i.ResetPassword.FirstByToken(db, token)
	if err != nil {
		return &models.ResetPasswords{}, services.NewResultStatus(404, err)
	}
	fmt.Println(foundResetPassword.ExpireAt, time.Now().Unix())

	if foundResetPassword.ExpireAt < time.Now().Unix() {
		return &models.ResetPasswords{}, services.NewResultStatus(400, fmt.Errorf("トークンの有効期限が切れています"))
	}

	return foundResetPassword, services.NewResultStatus(200, nil)
}

func (i *ResetPasswordInteractor) Create(email string) (*models.ResetPasswords, *services.ResultStatus) {

	db := i.DB.Begin()

	account, err := i.Account.FirstByEmail(db, email)
	if err != nil {
		db.Rollback()
		return &models.ResetPasswords{}, services.NewResultStatus(404, fmt.Errorf("このメールアドレスは登録されていません"))
	}

	user, err := i.User.FirstByAccountID(db, account.ID)
	if err != nil {
		db.Rollback()
		return &models.ResetPasswords{}, services.NewResultStatus(404, err)
	}

	resetPassword, err := i.ResetPassword.Create(db, &models.ResetPasswords{
		UserID:    user.ID,
		Token:     utils.RandomToken(25),
		ResetKey:  utils.RandomString(10),
		ExpireAt:  time.Now().AddDate(0, 0, 1).Unix(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		db.Rollback()
		return &models.ResetPasswords{}, services.NewResultStatus(404, err)
	}

	db.Commit()

	return resetPassword, services.NewResultStatus(200, nil)
}

func (i *ResetPasswordInteractor) Save(resetKey, password string) (*models.Users, *services.ResultStatus) {

	db := i.DB.Begin()

	if password == "" {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(404, fmt.Errorf("パスワードは必須です"))
	}

	foundResetPassword, err := i.ResetPassword.FirstByResetKey(db, resetKey)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(404, err)
	}

	if foundResetPassword.ExpireAt < time.Now().Unix() {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(400, err)
	}

	foundUser, err := i.User.FindByID(db, foundResetPassword.UserID)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(404, err)
	}

	foundAccount, err := i.Account.FindByID(db, foundUser.AccountID)

	foundAccount.Password, err = utils.GenerateFromPassword(password)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(400, err)
	}

	_, err = i.Account.Save(db, foundAccount)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(404, err)
	}

	foundResetPassword.ExpireAt = time.Now().Unix()

	_, err = i.ResetPassword.Save(db, foundResetPassword)
	if err != nil {
		db.Rollback()
		return &models.Users{}, services.NewResultStatus(400, err)
	}

	db.Commit()

	return foundUser, services.NewResultStatus(200, nil)
}
