package libjson

import (
	"encoding/json"
	"errors"
	customErrors "github/JosacabDev/api-sqlite/pkg/errors"
	"net/http"
)

func Encode(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func EncodeError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := map[string]string{"error": message}
	Encode(w, statusCode, errorResponse)
}

func EncodeCustomError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, &customErrors.HTTPError{}):
		httpError := err.(*customErrors.HTTPError)
		EncodeError(w, httpError.StatusCode, httpError.Message)
		return
	case errors.Is(err, &customErrors.AppError{}):
		appError := err.(*customErrors.AppError)
		if appError.ErrorCode == customErrors.ErrorCodeNotFound {
			EncodeError(w, http.StatusNotFound, appError.Message)
		} else {
			EncodeError(w, http.StatusInternalServerError, appError.Message)
		}
		return

	default:
		// For any other error, return a generic internal server error
		EncodeError(w, http.StatusInternalServerError, err.Error())
	}
}

func EncodeCreated(w http.ResponseWriter, data interface{}) {
	Encode(w, http.StatusCreated, data)
}

func EncodeOk(w http.ResponseWriter, data interface{}) {
	Encode(w, http.StatusOK, data)
}

func EncodeNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
