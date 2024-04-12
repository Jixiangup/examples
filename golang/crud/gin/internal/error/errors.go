package error

import (
	"errors"
	"net/http"
)

type Type int

const (
	Validation Type = iota + 1
	UnknownMistake
	Unauthorized
)

type ApplicationError struct {
	Type    Type
	Message string
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func NewValidation(message string) *ApplicationError {
	return &ApplicationError{
		Type:    Validation,
		Message: message,
	}
}

func NewUnknownMistake(message string) *ApplicationError {
	return &ApplicationError{
		Type:    UnknownMistake,
		Message: message,
	}
}
func NewUnauthorized(message string) *ApplicationError {
	return &ApplicationError{
		Type:    Unauthorized,
		Message: message,
	}
}

func GetStatus(err error) int {
	var appError *ApplicationError
	if errors.As(err, &appError) {
		switch appError.Type {
		// 参数异常
		case Validation:
			return http.StatusBadRequest
		case UnknownMistake:
			return http.StatusInternalServerError
		case Unauthorized:
			return http.StatusUnauthorized
		}
	}
	return http.StatusInternalServerError
}
