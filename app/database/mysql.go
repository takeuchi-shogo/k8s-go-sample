package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB() *DB {
	return &DB{
		Connection: connection("", "", "", ""),
	}
}

func connection(host, username, password, dbName string) *gorm.DB {
	conn, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return conn
}
