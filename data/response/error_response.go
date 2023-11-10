package response

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors"`
}

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}