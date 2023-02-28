package controllers

import (
	"context"
	"errors"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type UserSearchFiltersGraphqlController struct {
	Authorize  interactor.AuthorizeInteractor
	Interactor interactor.UserSearchFilterInteractor
}

func NewUserSearchFilterGraphqlController(db repositories.DB, jwt gateways.Jwt) *UserSearchFiltersGraphqlController {
	return &UserSearchFiltersGraphqlController{
		Authorize: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
		Interactor: interactor.UserSearchFilterInteractor{
			DB:               &repositories.DBRepository{DB: db},
			UserSearchFilter: &repositories.UserSearchFilterRepository{},
		},
	}
}

func (controller *UserSearchFiltersGraphqlController) Get(ctx context.Context) (*models.UserSearchFilters, error) {
	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token := gc.GetHeader("authorization")
	if token == "" {
		return nil, helpers.GraphQLErrorResponse(ctx, errors.New("ログインしてください"), 401)
	}

	userID, res := controller.Authorize.Verify(token)
	if res.Error != nil {
		return &models.UserSearchFilters{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	filter, res := controller.Interactor.Get(userID)
	if res.Error != nil {
		return &models.UserSearchFilters{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	return filter, nil
}

func (controller *UserSearchFiltersGraphqlController) Post(ctx context.Context, f *models.UserSearchFilters) (*models.UserSearchFilters, error) {
	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token := gc.GetHeader("authorization")
	if token == "" {
		return nil, helpers.GraphQLErrorResponse(ctx, errors.New("ログインしてください"), 401)
	}

	userID, res := controller.Authorize.Verify(token)
	if res.Error != nil {
		return &models.UserSearchFilters{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	newFilter, res := controller.Interactor.Create(&models.UserSearchFilters{
		UserID:   userID,
		Gender:   f.Gender,
		Location: f.Location,
	})
	if res.Error != nil {
		return &models.UserSearchFilters{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return newFilter, nil
}

func (controller *UserSearchFiltersGraphqlController) Patch(ctx context.Context, filter *models.UserSearchFilters) (*models.UserSearchFilters, error) {
	updatedFilter, res := controller.Interactor.Save(filter)
	if res.Error != nil {
		return &models.UserSearchFilters{}, helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}
	return updatedFilter, nil
}
