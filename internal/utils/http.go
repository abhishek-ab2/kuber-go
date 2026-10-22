package utils

import "net/http"

type HttpResponse struct {
	Message string `json:"msg,omitempty"`
	Data    any    `json:"data,omitempty"`
	Success bool   `json:"success"`
}

func GeneralResponse(success bool, msg string, data any, status int) (int, HttpResponse) {
	if status == 0 {
		status = http.StatusOK
	}
	return status, HttpResponse{
		Success: success,
		Message: msg,
		Data:    data,
	}
}
