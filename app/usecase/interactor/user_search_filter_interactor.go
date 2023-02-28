package interactor

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type UserSearchFilterInteractor struct {
	DB               repository.DBRepository
	UserSearchFilter repository.UserSearchFilterRepository
}

func (interactor *UserSearchFilterInteractor) Get(userID int) (*models.UserSearchFilters, *services.ResultStatus) {
	db := interactor.DB.Connect()

	filter, err := interactor.UserSearchFilter.FirstByUserID(db, userID)
	if err != nil {
		return &models.UserSearchFilters{}, services.NewResultStatus(400, err)
	}
	return filter, services.NewResultStatus(200, nil)
}

func (interactor *UserSearchFilterInteractor) Create(filter *models.UserSearchFilters) (*models.UserSearchFilters, *services.ResultStatus) {
	db := interactor.DB.Connect()
	newFilter, err := interactor.UserSearchFilter.Create(db, filter)
	if err != nil {
		return &models.UserSearchFilters{}, services.NewResultStatus(400, err)
	}
	return newFilter, services.NewResultStatus(200, nil)
}

func (interactor *UserSearchFilterInteractor) Save(filter *models.UserSearchFilters) (*models.UserSearchFilters, *services.ResultStatus) {
	db := interactor.DB.Connect()

	f, err := interactor.UserSearchFilter.FindByID(db, filter.ID)
	if err != nil {
		return &models.UserSearchFilters{}, services.NewResultStatus(400, err)
	}

	f = filter

	updatedFilter, err := interactor.UserSearchFilter.Save(db, f)
	if err != nil {
		return &models.UserSearchFilters{}, services.NewResultStatus(400, err)
	}
	return updatedFilter, services.NewResultStatus(200, nil)
}
