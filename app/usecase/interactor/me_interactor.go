package interactor

import (
	"fmt"
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type MeInteractor struct {
	DB          repository.DBRepository
	Like        repository.LikeRepository
	User        repository.UserRepository
	UserProfile repository.UserProfileRepository

	UserProfilePresenter presenter.UserProfilePresenter
}

func (i *MeInteractor) Get(id int) (*models.ResponseUsers, *services.ResultStatus) {
	db := i.DB.Connect()

	user, err := i.User.FindByID(db, id)
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

	profile, _ := i.UserProfile.FindByUserID(db, user.ID)
	fmt.Println(profile)

	builtProfile := i.UserProfilePresenter.ResponseUserProfile(profile)

	ru.UserProfile = builtProfile
	return ru, services.NewResultStatus(http.StatusOK, nil)
}
