package rest

import (
	"fmt"
	"net/http"
)

var ErrInternalServerError = NewHTTPError(http.StatusInternalServerError)

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code    int
	Message string
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int) *HTTPError {
	return &HTTPError{Code: code, Message: http.StatusText(code)}
}

// Error makes it compatible with `error` interface.
func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}
