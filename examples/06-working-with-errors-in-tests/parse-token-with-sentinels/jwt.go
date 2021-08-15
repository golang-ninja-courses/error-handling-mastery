package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	supportedTokenType = "JWT"
	supportedAlgo      = "HS256"
)

var byteDot = []byte(".")

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

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Token struct {
	Email     string   `json:"email"`
	Subject   string   `json:"subject"`
	Scopes    []string `json:"scopes"`
	ExpiredAt int64    `json:"expired_at"`
}

// ParseToken парсит и валидирует токен jwt, проверяя, что он подписан
// алгоритмом HMAC SHA256 с использованием ключа secret.
func ParseToken(jwt, secret []byte) (Token, error) {
	if len(jwt) == 0 {
		return Token{}, ErrEmptyJWT
	}

	parts := bytes.Split(jwt, byteDot)
	if len(parts) != 3 {
		return Token{}, ErrInvalidTokenFormat
	}

	headerData, payloadData, signData := parts[0], parts[1], parts[2]

	h, err := parseHeader(headerData)
	if err != nil {
		return Token{}, fmt.Errorf("%w: %v", ErrInvalidHeaderEncoding, err)
	}

	if h.Typ != supportedTokenType {
		return Token{}, fmt.Errorf("%w: %q", ErrUnsupportedTokenType, h.Typ)
	}

	if err := verifySignature(
		h.Alg,
		bytes.Join([][]byte{parts[0], parts[1]}, byteDot),
		signData,
		secret,
	); err != nil {
		return Token{}, fmt.Errorf("verify signature: %w", err)
	}

	t, err := parsePayload(payloadData)
	if err != nil {
		return Token{}, fmt.Errorf("%w: %v", ErrInvalidPayloadEncoding, err)
	}

	if time.Unix(t.ExpiredAt, 0).Before(time.Now()) {
		return Token{}, ErrExpiredToken
	}

	return t, nil
}

func parseHeader(data []byte) (Header, error) {
	d := json.NewDecoder(
		base64.NewDecoder(base64.RawURLEncoding, bytes.NewReader(data)),
	)

	var header Header
	return header, d.Decode(&header)
}

func verifySignature(algo string, unsignedData, receivedSignature, secret []byte) error {
	if algo != supportedAlgo {
		return fmt.Errorf("%w: %q", ErrUnsupportedSigningAlgo, algo)
	}

	h := hmac.New(sha256.New, secret)
	h.Write(unsignedData)

	enc := base64.RawURLEncoding
	sign := make([]byte, enc.EncodedLen(h.Size()))
	enc.Encode(sign, h.Sum(nil))

	if !hmac.Equal(sign, receivedSignature) {
		return fmt.Errorf("%w: %q vs expected %q", ErrInvalidSignature, receivedSignature, sign)
	}

	return nil
}

func parsePayload(data []byte) (Token, error) {
	d := json.NewDecoder(
		base64.NewDecoder(base64.RawURLEncoding, bytes.NewReader(data)),
	)

	var token Token
	return token, d.Decode(&token)
}
