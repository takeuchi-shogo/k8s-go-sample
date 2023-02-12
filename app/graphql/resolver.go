package graphql

import "github.com/takeuchi-shogo/k8s-go-sample/domain/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// todos []*model.Todo
	account *models.Accounts
	block   *models.Blocks
	report  *models.Reports
	users   []*models.Users
	user    *models.Users
}
