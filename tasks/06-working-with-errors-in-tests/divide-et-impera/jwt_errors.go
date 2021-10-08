package jwt

import "errors"

var (
	ErrInvalidHeaderEncoding  = errors.New("invalid header encoding")
	ErrInvalidPayloadEncoding = errors.New("invalid payload encoding")

	ErrInvalidBase64 = errors.New("invalid base64")
	ErrInvalidJSON   = errors.New("invalid json")
)
