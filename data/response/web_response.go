package response

type WebResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}