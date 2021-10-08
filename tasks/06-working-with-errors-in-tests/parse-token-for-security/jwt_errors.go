package jwt

import "errors"

var (
	ErrEmptyJWT               = errors.New("empty jwt data")
	ErrInvalidTokenFormat     = errors.New("invalid token format: 'header.payload.signature' was expected")
	ErrInvalidHeaderEncoding  = errors.New("invalid header encoding")
	ErrUnsupportedTokenType   = errors.New("unsupported token type")
	ErrUnsupportedSigningAlgo = errors.New("unsupported the signing algo")
	ErrInvalidSignature       = errors.New("invalid signature")
	ErrInvalidPayloadEncoding = errors.New("invalid payload encoding")
	ErrExpiredToken           = errors.New("token was expired")
)
