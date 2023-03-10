package interactor

import (
	"net/http"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type UserProfileInteractor struct {
	DBRepository          repository.DBRepository
	UserProfileRepository repository.UserProfileRepository
}

func (interactor *UserProfileInteractor) Get(id int) (*models.UserProfiles, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	profile, err := interactor.UserProfileRepository.FindByID(db, id)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return profile, nil
}

func (interactor *UserProfileInteractor) Create(p *models.UserProfiles) (*models.UserProfiles, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	profile, err := interactor.UserProfileRepository.Create(db, p)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return profile, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *UserProfileInteractor) Save(up *models.UserProfiles) (*models.UserProfiles, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	foundProfile, err := interactor.UserProfileRepository.FindByID(db, up.ID)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}

	foundProfile.BodyTypeID = up.BodyTypeID

	profile, err := interactor.UserProfileRepository.Save(db, foundProfile)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return profile, services.NewResultStatus(http.StatusOK, nil)
}
