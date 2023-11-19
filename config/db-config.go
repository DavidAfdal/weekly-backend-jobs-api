package config

import (
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/jobfinder_db?charset=utf8&parseTime=True&loc=Local"))

	helper.ErrorPanic(err)

	db.AutoMigrate(&models.Job{})
	db.AutoMigrate(&models.Apllier{})


    
	return db
}