package models

import (
	"time"
)

type Job struct {
	Id           int    `gorm:"type:int;primaryKey" json:"id" `
	Title        string `json:"title"`
	Description  string `json:"description,omitempty" `
	Company      string `json:"company,omitempty"`
	ImageCompany string `json:"image,omitempty"`
	Category     string `json:"category,omitempty"`
	Status       string `json:"status,omitempty"`
	Location     string `json:"location,omitempty"`
	Salary       int64  `json:"salary,omitempty"`
	UserId         string `json:"userId,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`

}