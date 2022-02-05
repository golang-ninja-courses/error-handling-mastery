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
		expectedEmail string
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
			name:          "invalid header encoding",
			jwt:           "XXXXbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			expectedErr:   ErrInvalidHeaderEncoding,
			expectedEmail: "john@gmail.com",
		},
		{
			name: "unsupported token type",
			// {"alg": "HS256", "typ":  "JWT123"}
			// {"subject": "1234567890", "email": "bob@gmail.com", "scopes": ["admin"], "expired_at": 4104586081}
			jwt:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVDEyMyJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiYm9iQGdtYWlsLmNvbSIsInNjb3BlcyI6WyJhZG1pbiJdLCJleHBpcmVkX2F0Ijo0MTA0NTg2MDgxfQ.9iwM6OKHXwqLa-jx67-RPtAiUHIOqAf0g-LGDirc-PI",
			expectedErr:   ErrUnsupportedTokenType,
			expectedEmail: "bob@gmail.com",
		},
		{
			// {"alg": "HS512", "typ":  "JWT"}
			name:          "unsupported signing algo",
			jwt:           "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.YeYcZt06wTfXo62dD8qNvkhbersq_DySM3jmfS0SdeexTCQPbDoqfzoV3HX023zufGIfuSpri5OnyYX39ABVxg",
			expectedErr:   ErrUnsupportedSigningAlgo,
			expectedEmail: "john@gmail.com",
		},
		{
			// signed by "secret123"
			name:          "invalid signature",
			jwt:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.EVWFGUfF7ZsJoz7TIXKV0SmJkc2VjYa6zniEIHDnPgk",
			expectedErr:   ErrInvalidSignature,
			expectedEmail: "john@gmail.com",
		},
		{
			name:        "invalid payload encoding",
			jwt:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XXXdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.OB1AerP__4yOYkpvoV9GL6CKjra_IYSFLRfSEk6biw8",
			expectedErr: ErrInvalidPayloadEncoding,
		},
		{
			// 1611600481
			name:          "expired token",
			jwt:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiYm9iQGdtYWlsLmNvbSIsInNjb3BlcyI6WyJhZG1pbiJdLCJleHBpcmVkX2F0IjoxNjExNjAwNDgxfQ._e4SPuxQ4CrQ1Q25_8Vi00tGEDgnN1Ib2HzkrrLd-38",
			expectedErr:   ErrExpiredToken,
			expectedEmail: "bob@gmail.com",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token, err := ParseToken([]byte(tt.jwt), []byte("secret"))

			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("want error %q, got %q", tt.expectedErr, err)
			}

			if tt.expectedEmail != "" {
				var e interface {
					Email() string
				}
				if !errors.As(err, &e) {
					t.Fatalf("error %T doesn't implement `Email() string` method", err)
				}

				if e.Email() != tt.expectedEmail {
					t.Fatalf("got email %q, want: %q", e.Email(), tt.expectedEmail)
				}
			}

			if !reflect.DeepEqual(token, tt.expectedToken) {
				t.Fatalf("got token %#v, want: %#v", token, tt.expectedToken)
			}
		})
	}
}
