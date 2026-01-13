package apperrors

import (
	"net/http"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewCustomError(message string, statuscode int) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statuscode,
	}
}

func NotFoundError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func BadException(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		StatusCode: http.StatusBadRequest,
	}
}

func InternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		StatusCode: http.StatusInternalServerError,
	}
}

func UnprocessableEntityError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		StatusCode: http.StatusUnprocessableEntity,
	}
} 

func ConflictError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		StatusCode: http.StatusConflict,
	}
}

func UnauthorizedException(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		StatusCode: http.StatusUnauthorized,
	}
}