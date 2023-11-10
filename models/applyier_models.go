package models

import (
	"time"
)

type Apllier struct {
	Id   string  `gorm:"type:int;primaryKey"`
	Name string 
	UserID string 
	Jobs []Job  `gorm:"many2many:applier_job;"`
	CreatedAt time.Time 
}