package presenters

import (
	"log"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
)

type UsersPresenter struct {
}

func NewUsersPresenter() presenter.UserPresenter {
	return &UsersPresenter{}
}

func (presenter *UsersPresenter) ResponseUser(user *models.Users) *models.Users {
	user.Name = "Mr." + user.Name
	log.Println(user)
	return user
}
