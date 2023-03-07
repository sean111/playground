package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://blog.canopas.com/golang-gorm-with-mysql-and-gin-ab876f406244

var DB *gorm.DB

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open()

	if err != nil {
		return err
	}
	return db, nil
}
