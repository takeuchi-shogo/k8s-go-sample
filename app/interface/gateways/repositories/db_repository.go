package repositories

import "gorm.io/gorm"

type DBRepository struct {
	DB DB
}

func (db *DBRepository) Connect() *gorm.DB {
	return db.Connect()
}
