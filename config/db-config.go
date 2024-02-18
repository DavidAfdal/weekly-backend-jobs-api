package config

import (
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDb() *gorm.DB {
	u, err := dburl.Parse("mysql://root:cfa2bCcBbgbeh311GHBCG46hdgCbCgg4@roundhouse.proxy.rlwy.net:44435/railway" + "?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
		helper.ErrorPanic(err)
    }
	db, err := gorm.Open(mysql.Open(u.DSN), &gorm.Config{})

	helper.ErrorPanic(err)

	db.AutoMigrate(&models.Job{})
	db.AutoMigrate(&models.Apllier{})


    
	return db
}