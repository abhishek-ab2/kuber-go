package internal

import (
	"fmt"
	"net/http"
)

const (
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeServerError  = "SERVER_ERROR"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden    = "FORBIDDEN"
)

type BaseError struct {
	Message string
	Code    string
	Data    interface{}
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%s: %s \n %v", e.Code, e.Message, e.Data)
}

type APIError struct {
	BaseError
	StatusCode int
}

func ValidationError(err string, message string, data interface{}) *APIError {
	code := err
	if code == "" {
		code = ErrCodeValidation
	}
	return &APIError{
		BaseError: BaseError{
			Message: message,
			Code:    code,
			Data:    data,
		},
		StatusCode: http.StatusBadRequest,
	}
}

func NotFoundError(message string, data interface{}) *APIError {
	return &APIError{
		BaseError: BaseError{
			Message: message,
			Code:    ErrCodeNotFound,
			Data:    data,
		},
		StatusCode: http.StatusNotFound,
	}
}

func ServerError(message string, data interface{}) *APIError {
	return &APIError{
		BaseError: BaseError{
			Message: message,
			Code:    ErrCodeServerError,
			Data:    data,
		},
		StatusCode: http.StatusInternalServerError,
	}
}

func UnauthorizedError(message string, data interface{}) *APIError {
	return &APIError{
		BaseError: BaseError{
			Message: message,
			Code:    ErrCodeUnauthorized,
			Data:    data,
		},
		StatusCode: http.StatusUnauthorized,
	}
}

func ForbiddenError(message string, data interface{}) *APIError {
	return &APIError{
		BaseError: BaseError{
			Message: message,
			Code:    ErrCodeForbidden,
			Data:    data,
		},
		StatusCode: http.StatusForbidden,
	}
}
