package jwt

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseToken_v2(t *testing.T) {
	// Для составления тест-кейсов был использован https://jwt.io/
	// с алгоритмом HS256 и "secret" в качестве ключа хеширования.
	cases := []struct {
		name            string
		jwt             string
		expectedErrText string
		expectedToken   Token
	}{
		{
			// {"alg": "HS256", "typ":  "JWT"}
			// {"subject": "1234567890", "email": "john@gmail.com", "scopes": ["admin"], "expired_at": 4104586081}
			name:            "absolutely valid token",
			jwt:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			expectedErrText: "",
			expectedToken: Token{
				Email:     "john@gmail.com",
				Subject:   "1234567890",
				Scopes:    []string{"admin"},
				ExpiredAt: 4104586081,
			},
		},
		{
			name:            "empty jwt",
			jwt:             "",
			expectedErrText: "empty jwt data",
		},
		{
			name:            "invalid token format",
			jwt:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0",
			expectedErrText: "invalid token format",
		},
		{
			name:            "invalid header encoding",
			jwt:             "XXXXbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			expectedErrText: "invalid character", // invalid character ']' looking for beginning of value
		},
		{
			name: "unsupported token type",
			// {"alg": "HS256", "typ":  "JWT123"}
			jwt:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVDEyMyJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.6CM6Ys-eI1GUM0L0FG0d2yavG29FDujcklMc0ipNhXA",
			expectedErrText: "unsupported token type",
		},
		{
			// {"alg": "HS512", "typ":  "JWT"}
			name:            "unsupported signing algo",
			jwt:             "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.YeYcZt06wTfXo62dD8qNvkhbersq_DySM3jmfS0SdeexTCQPbDoqfzoV3HX023zufGIfuSpri5OnyYX39ABVxg",
			expectedErrText: "unsupported the signing algo",
		},
		{
			// signed by "secret123"
			name:            "invalid signature",
			jwt:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.EVWFGUfF7ZsJoz7TIXKV0SmJkc2VjYa6zniEIHDnPgk",
			expectedErrText: "invalid signature",
		},
		{
			name:            "invalid payload encoding",
			jwt:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XXXdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.OB1AerP__4yOYkpvoV9GL6CKjra_IYSFLRfSEk6biw8",
			expectedErrText: "invalid character", // invalid character ']' looking for beginning of value
		},
		{
			// 1611600481
			name:            "expired token",
			jwt:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6MTYxMTYwMDQ4MX0.MIVf6keNZGkoJPCajltFM7JNHPk6RXwmYxJbR_8_TE4",
			expectedErrText: "token was expired",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token, err := ParseToken([]byte(tt.jwt), []byte("secret"))

			if tt.expectedErrText == "" && (err != nil) {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.expectedErrText != "" {
				if err == nil {
					t.Fatalf("error with %q was expected: got <nil>", tt.expectedErrText)
				}

				if !strings.Contains(err.Error(), tt.expectedErrText) {
					t.Fatalf("error with %q was expected, got %q", tt.expectedErrText, err.Error())
				}
			}

			if !reflect.DeepEqual(token, tt.expectedToken) {
				t.Fatalf("got token %#v, want: %#v", token, tt.expectedToken)
			}
		})
	}
}
