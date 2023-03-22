package controllers

import (
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

type TweetsGraphqlController struct {
	Interactor interactor.TweetInteractor
}

func NewTweetsGraphqlController(db repositories.DB) *TweetsGraphqlController {
	return &TweetsGraphqlController{
		Interactor: interactor.TweetInteractor{
			DB:    &repositories.DBRepository{DB: db},
			Tweet: &repositories.TweetRepository{},
		},
	}
}
