package graph

import "github.com/takeuchi-shogo/k8s-go-sample/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
	user  *model.User
}
