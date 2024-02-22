// utils/errors.go
package utils

import (
	"net/http"
)

// Error represents a custom error type.
type Error struct {
	Message string
	Status  int // HTTP status code
}

// NewError creates a new Error with the specified status code and message.
func NewError(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}

// BadRequestError creates a new Error with a BadRequest status code.
func BadRequestError(message string) *Error {
	return NewError(http.StatusBadRequest, message)
}

// InternalServerError creates a new Error with an InternalServerError status code.
func InternalServerError(message string) *Error {
	return NewError(http.StatusInternalServerError, message)
}

// NotFoundError creates a new Error with a NotFound status code.
func NotFoundError(message string) *Error {
	return NewError(http.StatusNotFound, message)
}

// Add more custom error types if needed...

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}
