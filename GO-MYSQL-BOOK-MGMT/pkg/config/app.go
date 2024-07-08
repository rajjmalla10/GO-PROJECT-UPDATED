package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3308)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db = d
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
