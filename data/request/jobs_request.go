package request

type CreateJobInput struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description,omitempty" binding:"required"`
	Company      string `json:"company,omitempty" binding:"required"`
	ImageCompany string `json:"image,omitempty" binding:"required"`
	Category     string `json:"category,omitempty" binding:"required"`
	Status       string `json:"status,omitempty" binding:"required"`
	Location     string `json:"location,omitempty" binding:"required"`
	Salary       int64  `json:"salary,omitempty" binding:"required"`
	UserId       string `json:"userId,omitempty" binding:"required"`
}

type UpdateJobInput struct {
	Id           int    `json:"id"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description,omitempty" binding:"required"`
	Company      string `json:"company,omitempty" binding:"required"`
	ImageCompany string `json:"image,omitempty" binding:"required"`
	Category     string `json:"category,omitempty" binding:"required"`
	Status       string `json:"status,omitempty" binding:"required"`
	Location     string `json:"location,omitempty" binding:"required"`
	Salary       int64  `json:"salary,omitempty" binding:"required"`
}
