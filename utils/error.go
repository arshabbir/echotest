package utils

type ApiError struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
}
