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
	LikeRepository             repository.LikeRepository
	UserRepository             repository.UserRepository
	UserProfileRepository      repository.UserProfileRepository
	UserPresenter              presenter.UserPresenter
	UserSearchFilterRepository repository.UserSearchFilterRepository
	VerifyEmailRepository      repository.VerifyEmailRepository

	UserProfilePresenter presenter.UserProfilePresenter
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
		profile, _ := interactor.UserProfileRepository.FindByUserID(db, user.ID)
		ru := &models.ResponseUsers{}
		ru.ID = user.ID
		ru.DisplayName = user.DisplayName
		ru.ScreenName = user.ScreenName
		ru.Gender = user.Gender

		builtProfile := interactor.UserProfilePresenter.ResponseUserProfile(profile)

		ru.UserProfile = builtProfile
		builtUser := interactor.UserPresenter.ResponseUser(ru)
		userEdge := &types.UserEdge{
			Cursor: getCursor(user),
			Node:   builtUser,
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

func (interactor *UserInteractor) Get(id, userID int) (*models.ResponseUsers, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	user, err := interactor.UserRepository.FindByID(db, userID)
	if err != nil {
		return &models.ResponseUsers{}, services.NewResultStatus(http.StatusNotFound, err)
	}

	ru := &models.ResponseUsers{
		ID:          user.ID,
		UUID:        user.UUID,
		DisplayName: user.DisplayName,
		ScreenName:  user.ScreenName,
		Age:         user.Age,
		Gender:      user.Gender,
		Location:    user.Location,
	}

	if _, err := interactor.LikeRepository.FindBySendUserIDAndReceiveUserID(db, id, userID); err == nil {
		ru.IsLiked = true
	}

	profile, _ := interactor.UserProfileRepository.FindByUserID(db, user.ID)

	builtProfile := interactor.UserProfilePresenter.ResponseUserProfile(profile)

	ru.UserProfile = builtProfile
	return ru, services.NewResultStatus(http.StatusOK, nil)
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
	// foundUser.ScreenName = user.ScreenName
	foundUser.Gender = user.Gender
	foundUser.Age = user.Age
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
