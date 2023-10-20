package models

import "time"

type Job struct {
	Id           int    `json:"job_id" gorm:"type:int;primary_key"`
	Title        string `json:"job_title"`
	Description  string `json:"job_description"`
	Company      string `json:"job_company"`
	ImageCompany string `json:"job_image"`
	Category     string `json:"job_category"`
	Status       string `json:"job_status"`
	Location     string `json:"job_location"`
	Salary       int64  `json:"job_salary"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
}