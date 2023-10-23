package request

type CreateJobInput struct {
	Title        string `json:"title"`
	Description  string `json:"description,omitempty" `
	Company      string `json:"company,omitempty"`
	ImageCompany string `json:"image,omitempty"`
	Category     string `json:"category,omitempty"`
	Status       string `json:"status,omitempty"`
	Location     string `json:"location,omitempty"`
	Salary       int64  `json:"job_salary,omitempty"`
}

type UpdateJobInput struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty" `
	Company      string `json:"company,omitempty"`
	ImageCompany string `json:"image,omitempty"`
	Category     string `json:"category,omitempty"`
	Status       string `json:"status,omitempty"`
	Location     string `json:"location,omitempty"`
	Salary       int64  `json:"job_salary,omitempty"`
}