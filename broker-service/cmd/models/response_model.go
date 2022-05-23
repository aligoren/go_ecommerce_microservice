package models

type ResponseModel struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      bool   `json:"error"`
	Data       any    `json:"data,omitempty"`
}
