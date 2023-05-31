package interactor

import (
	"errors"
	"net/http"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/gateway"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
	"github.com/takeuchi-shogo/k8s-go-sample/utils"
	"gorm.io/gorm"
)

type AccountInteractor struct {
	AccountRepository          repository.AccountRepository
	DBRepository               repository.DBRepository
	Jwt                        gateway.JwtGateway
	UserRepository             repository.UserRepository
	UserProfileRepository      repository.UserProfileRepository
	UserSearchFilterRepository repository.UserSearchFilterRepository
	VerifyEmailRepository      repository.VerifyEmailRepository

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

func (interactor *AccountInteractor) setVelue(db *gorm.DB) string {
	var value string

	for {
		value = utils.RandomScreenName()
		if _, err := interactor.UserRepository.FirstByScreenName(db, value); err != nil {
			// 見つからなければfor文を抜ける
			break
		}
	}

	return value
}

func (interactor *AccountInteractor) Signup(account *models.Accounts, user *models.Users) (*models.Users, string, *services.ResultStatus) {
	db := interactor.DBRepository.Begin()

	account.PhoneNumber = "000000"
	account.Password, _ = utils.GenerateFromPassword(account.Password)
	account.LoginStatus = "new_account"
	account.AccessLevel = 1
	currentTime := time.Now().Unix()
	account.CreatedAt = currentTime
	account.UpdatedAt = currentTime

	createdAccount, err := interactor.AccountRepository.Create(db, account)
	if err != nil {
		db.Rollback()
		return &models.Users{}, "", services.NewResultStatus(http.StatusBadRequest, err)
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
		return &models.Users{}, "", services.NewResultStatus(http.StatusBadRequest, err)
	}
	var gender string = ""

	switch createdUser.Gender {
	case "M":
		gender = "F"
	case "F":
		gender = "M"
	}

	filter := &models.UserSearchFilters{
		UserID:    createdUser.ID,
		Gender:    &gender,
		Location:  nil,
		CreatedAt: services.SetNowDate(),
		UpdatedAt: services.SetNowDate(),
	}

	if _, err = interactor.UserSearchFilterRepository.Create(db, filter); err != nil {
		db.Rollback()
		return &models.Users{}, "", services.NewResultStatus(http.StatusBadRequest, err)
	}

	if _, err := interactor.UserProfileRepository.Create(db, &models.UserProfiles{
		UserID:    createdUser.ID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}); err != nil {
		db.Rollback()
		return &models.Users{}, "", services.NewResultStatus(http.StatusBadRequest, err)
	}

	jwtToken := interactor.Jwt.CreateToken(createdUser.ID)

	db.Commit()
	return createdUser, jwtToken, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *AccountInteractor) Save(userID int, a *types.UpdateAccounts) (*models.Accounts, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	foundUser, err := interactor.UserRepository.FindByID(db, userID)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	foundAccount, err := interactor.AccountRepository.FindByID(db, foundUser.AccountID)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	if a.ID != nil {
		id := *a.ID
		foundAccount.ID = id
	}
	if a.PhoneNumber != nil {
		phone := *a.PhoneNumber
		foundAccount.PhoneNumber = phone
	}
	if a.Email != nil {
		email := *a.Email
		foundAccount.Email = email
	}
	foundAccount.Password, err = generatePassword(foundAccount.Password, a.CurrentPassword, a.NewPassword)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	account, err := interactor.AccountRepository.Save(db, foundAccount)
	if err != nil {
		return &models.Accounts{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return account, services.NewResultStatus(http.StatusOK, nil)
}

func generatePassword(oldPassword string, currentPassword, newPassword *string) (string, error) {
	// NULL or undefind、nilの確認
	if currentPassword != nil {
		if *currentPassword != "" {
			pass := *currentPassword
			if err := utils.CheckPassword(pass, oldPassword); err != nil {
				return "", err
			}
			if newPassword != nil {
				if *newPassword != "" {
					return utils.GenerateFromPassword(*newPassword)
				} else {
					return "", errors.New("新しいパスワードを入力してください")
				}
			}
		} else {
			return "", errors.New("現在のパスワードを入力してください")
		}
	}

	return oldPassword, nil
}
