package interactor

import (
	"time"

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

	f.Gender = filter.Gender
	f.Location = filter.Location
	f.Purpose = filter.Purpose
	f.HasIntroduction = filter.HasIntroduction
	f.HeightID = filter.HeightID
	f.BodyTypeID = filter.BodyTypeID
	f.BloodTypeID = filter.BloodTypeID
	// f.ResidenceCountryID = filter.ResidenceCountryID
	// f.ResidenceStateID = filter.ResidenceStateID
	// f.HometownCountryID = filter.HometownCountryID
	// f.HometownStateID = filter.HometownStateID
	f.OccupationID = filter.OccupationID
	f.EducationID = filter.EducationID
	f.AnnualIncomeID = filter.AnnualIncomeID
	f.SmokingID = filter.SmokingID
	f.DrinkingID = filter.DrinkingID

	f.UpdatedAt = time.Now().Unix()

	updatedFilter, err := interactor.UserSearchFilter.Save(db, f)
	if err != nil {
		return &models.UserSearchFilters{}, services.NewResultStatus(400, err)
	}
	return updatedFilter, services.NewResultStatus(200, nil)
}
