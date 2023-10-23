package error

type RequestError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}