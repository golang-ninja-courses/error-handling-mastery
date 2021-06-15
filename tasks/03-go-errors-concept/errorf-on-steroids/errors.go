package errs

import (
	"errors"
	"fmt"
)

type SteroidError struct {
	message string
	errs    []error
}

func NewSteroidError(message string, errs []error) *SteroidError {
	return &SteroidError{
		message: message,
		errs:    errs,
	}
}

func (e *SteroidError) Error() string {
	return e.message
}

func (e *SteroidError) Is(target error) bool {
	for _, err := range e.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (e *SteroidError) As(target interface{}) bool {
	for _, err := range e.errs {
		if errors.As(err, &target) {
			return true
		}
	}
	return false
}

func Errorf(format string, args ...interface{}) error {
	message := fmt.Errorf(format, args...).Error()

	var errs []error
	for _, arg := range args {
		if err, ok := arg.(error); ok {
			errs = append(errs, err)
		}
	}

	return NewSteroidError(message, errs)
}
