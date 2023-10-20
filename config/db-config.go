package config

import (
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:LRI2znbbwgHKAxvBAyGn@tcp(containers-us-west-162.railway.app:5789)/railway?charset=utf8&parseTime=True&loc=Local"))

	helper.ErrorPanic(err)

	db.Table("jobs").AutoMigrate(&models.Job{})
    
	return db
}