package config

import (
	"fmt"
	"os"
	"weekly/go/gin/helper"
	"weekly/go/gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDb() *gorm.DB {
    db_url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	fmt.Println(os.Getenv("PGHOST"))

	helper.ErrorPanic(err)

	db.AutoMigrate(&models.Job{})
	db.AutoMigrate(&models.Apllier{})


    
	return db
}