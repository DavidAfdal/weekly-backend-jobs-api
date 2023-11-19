package models

import (
	"time"
)

type Apllier struct {
	Id   string  `gorm:"type:int;primaryKey" json:"id"`
	UserID string `json:"userId"`
	Jobs []Job  `gorm:"many2many:applier_job;" json:"jobs"`
	CreatedAt time.Time `json:"createdAt"`
}