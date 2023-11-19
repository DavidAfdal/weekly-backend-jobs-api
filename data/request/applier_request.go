package request

type ApplierRequest struct {
	UserId string `json:"user_id" binding:"required"`
	JobId  int    `json:"job_id" binding:"required"`
}