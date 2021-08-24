package repository

import (
	"log"
	"user-service/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB -> connection to Database
func DB() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	db.AutoMigrate(&model.User{})
	return db

}
