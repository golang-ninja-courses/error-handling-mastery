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
		return Token{}, errors.New("empty jwt data")
	}

	parts := bytes.Split(jwt, byteDot)
	if len(parts) != 3 {
		return Token{}, errors.New("invalid token format: 'header.payload.signature' was expected")
	}

	headerData, payloadData, signData := parts[0], parts[1], parts[2]

	h, err := parseHeader(headerData)
	if err != nil {
		return Token{}, err
	}

	if h.Typ != supportedTokenType {
		return Token{}, fmt.Errorf("unsupported token type: %q", h.Typ)
	}

	if err := verifySignature(
		h.Alg,
		bytes.Join([][]byte{parts[0], parts[1]}, byteDot),
		signData,
		secret,
	); err != nil {
		return Token{}, err
	}

	t, err := parsePayload(payloadData)
	if err != nil {
		return Token{}, err
	}

	if time.Unix(t.ExpiredAt, 0).Before(time.Now()) {
		return Token{}, errors.New("token was expired")
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
		return fmt.Errorf("unsupported the signing algo: %q", algo)
	}

	h := hmac.New(sha256.New, secret)
	h.Write(unsignedData)

	enc := base64.RawURLEncoding
	sign := make([]byte, enc.EncodedLen(h.Size()))
	enc.Encode(sign, h.Sum(nil))

	if !hmac.Equal(sign, receivedSignature) {
		return fmt.Errorf("invalid signature: %q vs expected %q", receivedSignature, sign)
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
