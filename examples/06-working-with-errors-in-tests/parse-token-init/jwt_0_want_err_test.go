package jwt

import (
	"reflect"
	"testing"
)

func TestParseToken(t *testing.T) {
	// Для составления тест-кейсов был использован https://jwt.io/
	// с алгоритмом HS256 и "secret" в качестве ключа хеширования.
	cases := []struct {
		name          string
		jwt           string
		wantErr       bool
		expectedToken Token
	}{
		{
			// {"alg": "HS256", "typ":  "JWT"}
			// {"subject": "1234567890", "email": "john@gmail.com", "scopes": ["admin"], "expired_at": 4104586081}
			name:    "absolutely valid token",
			jwt:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			wantErr: false,
			expectedToken: Token{
				Email:     "john@gmail.com",
				Subject:   "1234567890",
				Scopes:    []string{"admin"},
				ExpiredAt: 4104586081,
			},
		},
		{
			name:    "empty jwt",
			jwt:     "",
			wantErr: true,
		},
		{
			name:    "invalid token format",
			jwt:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0",
			wantErr: true,
		},
		{
			name:    "invalid header encoding",
			jwt:     "XXXXbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			wantErr: true,
		},
		{
			name: "unsupported token type",
			// {"alg": "HS256", "typ":  "JWT123"}
			jwt:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVDEyMyJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.6CM6Ys-eI1GUM0L0FG0d2yavG29FDujcklMc0ipNhXA",
			wantErr: true,
		},
		{
			// {"alg": "HS512", "typ":  "JWT"}
			name:    "unsupported signing algo",
			jwt:     "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVDEyMyJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.ryWY506Gynnr-WhTZYBJiXLtKnSFRiPgXBRiFVU_PJIvMwNAv0-p_dHcPF_x8JlwDxVqAYymogLFZqDFjRbp7w",
			wantErr: true,
		},
		{
			// signed by "secret123"
			name:    "invalid signature",
			jwt:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVDEyMyJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.pjWE6Zr1UzHYqMKvlywOgeBhQMGSWt_1xEIV30H2jxk",
			wantErr: true,
		},
		{
			name:    "invalid payload encoding",
			jwt:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XXXdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0.N7pFHBeew0mKBz4ULkim20QYbp7tcizR7Chdn4l32w8",
			wantErr: true,
		},
		{
			// 1611600481
			name:    "expired token",
			jwt:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6MTYxMTYwMDQ4MX0.MIVf6keNZGkoJPCajltFM7JNHPk6RXwmYxJbR_8_TE4",
			wantErr: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token, err := ParseToken([]byte(tt.jwt), []byte("secret"))

			if tt.wantErr && err == nil {
				t.Fatalf("error was expected: got <nil>")
			}

			if !tt.wantErr && (err != nil) {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(token, tt.expectedToken) {
				t.Fatalf("got token %#v, want: %#v", token, tt.expectedToken)
			}
		})
	}
}
