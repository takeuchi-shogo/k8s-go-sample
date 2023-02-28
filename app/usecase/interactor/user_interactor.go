package interactor

import (
	"net/http"
	"strconv"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type UserInteractor struct {
	AccountRepository          repository.AccountRepository
	DBRepository               repository.DBRepository
	UserRepository             repository.UserRepository
	UserPresenter              presenter.UserPresenter
	UserSearchFilterRepository repository.UserSearchFilterRepository
	VerifyEmailRepository      repository.VerifyEmailRepository
}

func getCursor(user *models.Users) string {
	return strconv.Itoa(user.ID)
}

func (interactor *UserInteractor) GetList(first int, after string, userID int) (*types.UserConnection, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	filter, _ := interactor.UserSearchFilterRepository.FirstByUserID(db, userID)

	users, err := interactor.UserRepository.FindByUserFilter(db, filter, first+1, after)
	if err != nil {
		return &types.UserConnection{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	edges := []*types.UserEdge{}

	for _, user := range users {
		userEdge := &types.UserEdge{
			Cursor: getCursor(user),
			Node:   user,
		}
		edges = append(edges, userEdge)
	}

	pageInfo := &types.PageInfo{
		HasNextPage:     len(users) > first,
		HasPreviousPage: after != "",
	}

	if pageInfo.HasNextPage {
		pageInfo.EndCursor = &users[len(users)-1].ScreenName
		edges = edges[:len(users)-1]
	}
	if len(users) > 0 {
		pageInfo.StartCursor = &users[0].ScreenName
	}

	return &types.UserConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}, services.NewResultStatus(http.StatusOK, nil)
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
