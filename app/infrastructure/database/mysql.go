package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB() *DB {
	return &DB{
		Connection: connection(
			"3306",      // host
			"root",      // user name
			"password",  // password
			"luka_test", // database name
		),
	}
}

func connection(host, username, password, dbName string) *gorm.DB {
	count := 0
	conn, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{})
	if err != nil {
		for {
			if err == nil {
				break
			}
			fmt.Print(".\n")
			time.Sleep(time.Second)
			count++
			// connection wait 10 seconds for database starting...
			if count > 10 {
				fmt.Print("database connection failed\n")
				panic(err.Error())
			}
			conn, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{})
		}
	}

	log.Print("database connection success\n")

	return conn
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
