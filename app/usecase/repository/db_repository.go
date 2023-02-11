package repository

import "gorm.io/gorm"

type DBRepository interface {
	Connect() *gorm.DB
}
