package errors

import "fmt"

// ErrorCode constants for application-specific errors
const (
	ErrorCodeNotFound = iota + 1
)

type AppError struct {
	Message   string `json:"message"`
	AppContex string `json:"app_context"`
	ErrorCode int    `json:"error_code"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("UserMessage - %s | ApplicationContext - %s.", e.Message, e.AppContex)
}

func NewAppError(message, appContext string, errorCode int) *AppError {
	return &AppError{
		Message:   message,
		AppContex: appContext,
		ErrorCode: errorCode,
	}
}
