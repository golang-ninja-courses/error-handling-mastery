package httperr

// Реализуй нас.
var (
	ErrStatusOK                  error
	ErrStatusBadRequest          error
	ErrStatusNotFound            error
	ErrStatusUnprocessableEntity error
	ErrStatusInternalServerError error
)

// Реализуй меня.
type HTTPError int
