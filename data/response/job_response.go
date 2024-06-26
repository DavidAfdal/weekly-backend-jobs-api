package response

import (
	"time"
)


type JobResponse struct {
	Id       int    `json:"id" `
	Title        string `json:"title"`
	Description  string `json:"description,omitempty" `
	Company      string `json:"company,omitempty"`
	ImageCompany string `json:"image,omitempty"`
	Category     string `json:"category,omitempty"`
	Status       string `json:"status,omitempty"`
	Location     string `json:"location,omitempty"`
	Salary       int64  `json:"salary,omitempty"`
	UserId       string `json:"userId,omitempty"`
	Appliers     interface{} `json:"appliers,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}

