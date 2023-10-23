package models

import "time"

type Job struct {
	Id           int    `json:"id" gorm:"type:int;primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Company      string `json:"company"`
	ImageCompany string `json:"image"`
	Category     string `json:"category"`
	Status       string `json:"status"`
	Location     string `json:"location"`
	Salary       int64  `json:"salary"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
}