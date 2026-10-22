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
	return GeneralResponse(false, msg, data, http.StatusInternalServerError)
}

func ValidationError(msg string, data any) (int, HttpResponse) {
	return GeneralResponse(false, msg, data, http.StatusBadRequest)
}

func NotFoundError(msg string) (int, HttpResponse) {
	return GeneralResponse(false, msg, nil, http.StatusNotFound)
}

func UnAuthorisedError(msg string) (int, HttpResponse) {
	return GeneralResponse(false, msg, nil, http.StatusUnauthorized)
}

func NotAllowedError(msg string) (int, HttpResponse) {
	return GeneralResponse(false, msg, nil, http.StatusForbidden)
}
