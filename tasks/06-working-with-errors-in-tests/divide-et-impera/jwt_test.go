package jwt

import (
	"errors"
	"reflect"
	"testing"
)

func Test_parseHeader(t *testing.T) {
	cases := []struct {
		name                string
		headerData          string
		expectedInternalErr error
		expectedHeader      Header
	}{
		{
			// {"alg": "HS256", "typ": "JWT"}
			name:           "absolutely valid header",
			headerData:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			expectedHeader: Header{Alg: "HS256", Typ: "JWT"},
		},
		{
			name:                "corrupted header base64",
			headerData:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCхз",
			expectedInternalErr: ErrInvalidBase64,
		},
		{
			name:                "invalid header base64 (with padding)",
			headerData:          "eyJhbGciOiAiSFMyNTYiLCAidHlwIjogIkoifQ==",
			expectedInternalErr: ErrInvalidBase64,
		},
		{
			// {"alg": "HS256", "typ": "JWT"
			name:                "invalid header json",
			headerData:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCI",
			expectedInternalErr: ErrInvalidJSON,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			header, err := parseHeader([]byte(tt.headerData))

			if tt.expectedInternalErr != nil {
				if !errors.Is(err, ErrInvalidHeaderEncoding) {
					t.Fatalf("error %v cannot be represented as ErrInvalidHeaderEncoding", err)
				}
			}

			if !errors.Is(err, tt.expectedInternalErr) {
				t.Fatalf("want internal error %v, got %v", tt.expectedInternalErr, err)
			}

			if !reflect.DeepEqual(header, tt.expectedHeader) {
				t.Fatalf("got header %#v, want: %#v", header, tt.expectedHeader)
			}
		})
	}
}

func Test_parsePayload(t *testing.T) {
	cases := []struct {
		name                string
		tokenData           string
		expectedInternalErr error
		expectedToken       Token
	}{
		{
			// {"subject": "1234567890", "email": "john@gmail.com", "scopes": ["admin"], "expired_at": 4104586081}
			name:                "absolutely valid payload",
			tokenData:           "eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4MX0",
			expectedInternalErr: nil,
			expectedToken: Token{
				Email:     "john@gmail.com",
				Subject:   "1234567890",
				Scopes:    []string{"admin"},
				ExpiredAt: 4104586081,
			},
		},
		{
			name:                "corrupted payload base64",
			tokenData:           "eyJzdWJqZWN0IjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiam9obkBnbWFpbC5jb20iLCJzY29wZXMiOlsiYWRtaW4iXSwiZXhwaXJlZF9hdCI6NDEwNDU4NjA4Mхз",
			expectedInternalErr: ErrInvalidBase64,
		},
		{
			name:                "invalid payload base64 (with padding)",
			tokenData:           "eyJzdWJqZWN0IjogIjEyMzQ1Njc4OTAiLCAiZW1haWwiOiAiam9obkBnbWFpbC5jb20iLCAic2NvcGVzIjogWyJhZG1pIl0sICJleHBpcmVkX2F0IjogNDEwNDU4NjA4MX0=",
			expectedInternalErr: ErrInvalidBase64,
		},
		{
			// {"subject": "1234567890", "email": "john@gmail.com", "scopes": ["admin"], "expired_at": 4104586081
			name:                "invalid payload json",
			tokenData:           "eyJzdWJqZWN0IjogIjEyMzQ1Njc4OTAiLCAiZW1haWwiOiAiam9obkBnbWFpbC5jb20iLCAic2NvcGVzIjogWyJhZG1pbiJdLCAiZXhwaXJlZF9hdCI6IDQxMDQ1ODYwODE",
			expectedInternalErr: ErrInvalidJSON,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token, err := parsePayload([]byte(tt.tokenData))

			if tt.expectedInternalErr != nil {
				if !errors.Is(err, ErrInvalidPayloadEncoding) {
					t.Fatalf("error %v cannot be represented as ErrInvalidPayloadEncoding", err)
				}
			}

			if !errors.Is(err, tt.expectedInternalErr) {
				t.Fatalf("want internal error %v, got %v", tt.expectedInternalErr, err)
			}

			if !reflect.DeepEqual(token, tt.expectedToken) {
				t.Fatalf("got token %#v, want: %#v", token, tt.expectedToken)
			}
		})
	}
}
