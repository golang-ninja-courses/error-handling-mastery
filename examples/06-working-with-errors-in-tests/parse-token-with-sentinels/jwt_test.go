package jwt

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseToken(t *testing.T) {
	// Для составления тест-кейсов был использован https://jwt.io/
	// с алгоритмом HS256 и "secret" в качестве ключа хеширования.
	cases := []struct {
		name          string
		jwt           string
		expectedErr   error
		expectedToken Token
	}{
		{
			// {"alg": "HS256", "typ":  "JWT"}
			// {"subject": "1234567890", "email": "john@gmail.com", "scopes": ["admin"], "expired_at": 4104586081}
			name:        "absolutely valid token",
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			expectedErr: nil,
			expectedToken: Token{
				Email:     "john@gmail.com",
				Subject:   "1234567890",
				Scopes:    []string{"admin"},
				ExpiredAt: 4104586081,
			},
		},
		{
			name:        "empty jwt",
			jwt:         "",
			expectedErr: ErrEmptyJWT,
		},
		{
			name:        "invalid token format",
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0",
			expectedErr: ErrInvalidTokenFormat,
		},
		{
			name:        "invalid header encoding",
			jwt:         "XXXXbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			expectedErr: ErrInvalidHeaderEncoding,
		},
		{
			name: "unsupported token type",
			// {"alg": "HS256", "typ":  "JWT123"}
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVDEyMyJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.6CM6Ys-eI1GUM0L0FG0d2yavG29FDujcklMc0ipNhXA",
			expectedErr: ErrUnsupportedTokenType,
		},
		{
			// {"alg": "HS512", "typ":  "JWT"}
			name:        "unsupported signing algo",
			jwt:         "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.YeYcZt06wTfXo62dD8qNvkhbersq_DySM3jmfS0SdeexTCQPbDoqfzoV3HX023zufGIfuSpri5OnyYX39ABVxg",
			expectedErr: ErrUnsupportedSigningAlgo,
		},
		{
			// signed by "secret123"
			name:        "invalid signature",
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.EVWFGUfF7ZsJoz7TIXKV0SmJkc2VjYa6zniEIHDnPgk",
			expectedErr: ErrInvalidSignature,
		},
		{
			name:        "invalid payload encoding",
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XXXdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.OB1AerP__4yOYkpvoV9GL6CKjra_IYSFLRfSEk6biw8",
			expectedErr: ErrInvalidPayloadEncoding,
		},
		{
			// 1611600481
			name:        "expired token",
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6MTYxMTYwMDQ4MX0.MIVf6keNZGkoJPCajltFM7JNHPk6RXwmYxJbR_8_TE4",
			expectedErr: ErrExpiredToken,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token, err := ParseToken([]byte(tt.jwt), []byte("secret"))

			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("want error %v, got %v", tt.expectedErr, err)
			}

			if !reflect.DeepEqual(token, tt.expectedToken) {
				t.Fatalf("got token %#v, want: %#v", token, tt.expectedToken)
			}
		})
	}
}
