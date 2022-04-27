package monad

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestM_Unpack(t *testing.T) {
	_, err := M{}.Unpack()
	assert.ErrorIs(t, err, ErrNoMonadValue)
}

func TestMonad(t *testing.T) {
	cases := []struct {
		name        string
		d           string
		expectedErr error
		expectedReq UserRegistrationRequest
	}{
		{
			name: "valid request",
			d:    `{"email":"bob@gmail.com","password":"bob"}`,
			expectedReq: UserRegistrationRequest{
				Email:    "bob@gmail.com",
				Password: "bob",
			},
		},
		{
			name:        "invalid json",
			d:           `{"email":"bob@gmail.com",`,
			expectedErr: errInvalidData,
		},
		{
			name:        "empty email",
			d:           `{"email":"","password":"bob"}`,
			expectedErr: errEmptyField,
		},
		{
			name:        "empty password",
			d:           `{"email":"","password":""}`,
			expectedErr: errEmptyField,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := ParseUserRequest(strings.NewReader(tt.d))
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedReq, req)
		})
	}
}

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ParseUserRequest вычитывает из входного ридера UserRegistrationRequest и валидирует его в monadic style.
func ParseUserRequest(r io.Reader) (UserRegistrationRequest, error) {
	res, err := Unit(r).
		Bind(unmarshalRequest).
		Bind(validateEmail).
		Bind(validatePassword).
		Unpack()
	if err != nil {
		return UserRegistrationRequest{}, err
	}
	return res.(UserRegistrationRequest), nil
}

var (
	errUnexpectedTypeInMonad = errors.New("unexpected type in monad")
	errInvalidData           = errors.New("invalid data")
	errEmptyField            = errors.New("field is empty")
)

func unmarshalRequest(v any) M {
	r, ok := v.(io.Reader)
	if !ok {
		return Err(fmt.Errorf("%w: %T", errUnexpectedTypeInMonad, v))
	}

	var req UserRegistrationRequest
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return Err(fmt.Errorf("%w: %v", errInvalidData, err))
	}
	return Unit(req)
}

func validateEmail(v any) M {
	req, ok := v.(UserRegistrationRequest)
	if !ok {
		return Err(fmt.Errorf("%w: %T", errUnexpectedTypeInMonad, v))
	}
	if req.Email == "" {
		return Err(fmt.Errorf("%w: Email", errEmptyField))
	}
	return Unit(v)
}

func validatePassword(v any) M {
	req, ok := v.(UserRegistrationRequest)
	if !ok {
		return Err(fmt.Errorf("%w: %T", errUnexpectedTypeInMonad, v))
	}
	if req.Password == "" {
		return Err(fmt.Errorf("%w: Password", errEmptyField))
	}
	return Unit(v)
}
