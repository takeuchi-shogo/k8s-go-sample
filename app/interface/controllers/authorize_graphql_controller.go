package controllers

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type AuthorizeGraphqlController struct {
	Interactor interactor.AuthorizeInteractor
}

func NewAuthorizeGraphqlController(db repositories.DB, jwt gateways.Jwt) *AuthorizeGraphqlController {
	return &AuthorizeGraphqlController{
		Interactor: interactor.AuthorizeInteractor{
			AccountRepository: &repositories.AccountRepository{},
			DBRepository:      &repositories.DBRepository{DB: db},
			Jwt:               &gateways.JwtGateway{Jwt: jwt},
			UserRepository:    &repositories.UserRepository{},
		},
	}
}

func (controller *AuthorizeGraphqlController) Login(ctx context.Context, email, password string) error {
	token, res := controller.Interactor.Login(email, password)
	if res.Error != nil {
		return helpers.GraphQLErrorResponse(ctx, res.Error, res.Code)
	}

	gc, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return err
	}
	gc.Header("Authorization", "Bearer "+token)
	log.Println("4")
	return nil
}
