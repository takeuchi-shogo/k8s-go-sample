package repository

import "gorm.io/gorm"

type DBRepository interface {
	Conn() *gorm.DB
}
