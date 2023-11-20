package models

import (
	"time"
)

type Apllier struct {
	Id   int  `gorm:"type:int;primaryKey" json:"id"`
	UserID string `json:"userId"`
	Jobs []Job  `gorm:"many2many:applier_job;constraint:OnDelete:CASCADE;" json:"jobs"`
	CreatedAt time.Time `json:"createdAt"`
}