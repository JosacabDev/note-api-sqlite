package errors

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("HTTP Error [%d] %s", err.StatusCode, err.Message)
}

func New(message string, statusCode int) *HTTPError {
	return &HTTPError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func BadRequestError(msg string) *HTTPError {
	return New(msg, http.StatusBadRequest)
}

func NotFoundError(record string) *HTTPError {
	return New(fmt.Sprintf("%s not found", record), http.StatusNotFound)
}

func InternalError(msg string) *HTTPError {
	return New(fmt.Sprintf("Error: %s", msg), http.StatusInternalServerError)
}
