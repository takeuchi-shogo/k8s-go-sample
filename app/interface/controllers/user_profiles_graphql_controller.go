package controllers

import (
	"context"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type UserProfilesGraphqlController struct {
	Interactor interactor.UserProfileInteractor
}

func NewUserProfilesGraphqlController(db repositories.DB) *UserProfilesGraphqlController {
	return &UserProfilesGraphqlController{
		Interactor: interactor.UserProfileInteractor{
			DBRepository:          &repositories.DBRepository{DB: db},
			UserProfileRepository: &repositories.UserProfileRepository{},
		},
	}
}

func (controller *UserProfilesGraphqlController) Get(ctx context.Context, userID int) (*models.UserProfiles, error) {

	userProfile, res := controller.Interactor.Get(userID)
	if res.Error != nil {
		return userProfile, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return userProfile, nil
}

func (controller *UserProfilesGraphqlController) Patch(ctx context.Context, userProfile *models.UserProfiles) (*models.UserProfiles, error) {

	updatedUserProfile, res := controller.Interactor.Save(userProfile)
	if res.Error != nil {
		return updatedUserProfile, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return updatedUserProfile, nil
}
