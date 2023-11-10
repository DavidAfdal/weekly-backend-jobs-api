package models

import (
	"time"
)

type Job struct {
	Id           int    `gorm:"type:int;primaryKey"`
	Title        string `gorm:"not null"`
	Description  string `gorm:"not null"`
	Company      string `gorm:"not null"`
	ImageCompany string `gorm:"not null"`
	Category     string `gorm:"not null"`
	Status       string `gorm:"not null"`
	Location     string `gorm:"not null"`
	Salary       int64  `gorm:"not null"`
	UserId       string `gorm:"not null"`
	CreatedAt    time.Time
}