package db

import "api/internal"

const (
	NO_DOCUMENTS_ERROR       = "No documents found"
	MULTIPLE_DOCUMENTS_ERROR = "Multiple documents found"
	CONNECTION_ERROR         = "Connection error"
	DECODE_ERROR             = "Decode error"
	COLLECTION_ERROR         = "Collection error"
	INDEX_ERROR              = "Index error"
	DATABASE_ERROR           = "Database error"
)

func DBError(err string, message string, data interface{}) *internal.APIError {
	code := err
	if code == "" {
		code = DATABASE_ERROR
	}
	return &internal.APIError{
		BaseError: internal.BaseError{
			Message: message,
			Code:    code,
			Data:    data,
		},
		StatusCode: 500,
	}
}
