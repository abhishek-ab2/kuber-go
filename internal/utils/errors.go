package utils

import (
	"fmt"
	"net/http"
)

type Error struct {
	Message string
	Code    string
	Data    any
}

func (err Error) Error() string {
	return fmt.Sprintf("[ERROR %s] %s \n Data: %v", err.Code, err.Message, err.Data)
}

func ServerError(msg string, data any) (int, HttpResponse) {
	return http.StatusInternalServerError, HttpResponse{
		Message: msg,
		Data:    data,
	}
}

func ValidationError(msg string, data any) (int, HttpResponse) {
	return http.StatusBadRequest, HttpResponse{
		Message: msg,
		Data:    data,
	}
}

func NotFoundError(msg string) (int, HttpResponse) {
	return http.StatusNotFound, HttpResponse{
		Message: msg,
		Data:    nil,
	}
}

func UnAuthorisedError(msg string) (int, HttpResponse) {
	return http.StatusUnauthorized, HttpResponse{
		Message: msg,
		Data:    nil,
	}
}

func NotAllowedError(msg string) (int, HttpResponse) {
	return http.StatusForbidden, HttpResponse{
		Message: msg,
		Data:    nil,
	}
}
