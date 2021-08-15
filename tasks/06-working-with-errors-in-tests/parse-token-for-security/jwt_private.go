package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const (
	supportedTokenType = "JWT"
	supportedAlgo      = "HS256"
)

var byteDot = []byte(".")

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
