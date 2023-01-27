package repository

import "github.com/takeuchi-shogo/k8s-go-sample/domain/models"

type UserRepository interface {
	FindByID(id int) (*models.Users, error)
	Create(*models.Users) (*models.Users, error)
}
