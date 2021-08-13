package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"go.uber.org/multierr"
)

type UserRegistrationRequest struct {
	err      error
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRegistrationRequest) Err() error {
	return u.err
}

func (u *UserRegistrationRequest) Unmarshal(r io.Reader) {
	u.err = multierr.Append(u.err, json.NewDecoder(r).Decode(u))
}

func (u *UserRegistrationRequest) ValidateEmail() {
	if u.Email == "" {
		u.err = multierr.Append(u.err, errors.New("empty email"))
	}
}

func (u *UserRegistrationRequest) ValidatePassword() {
	if u.Password == "" {
		u.err = multierr.Append(u.err, errors.New("empty password"))
	}
}

func main() {
	var req UserRegistrationRequest
	req.Unmarshal(strings.NewReader(`{"email":"bob@gmail.com","password":"`))
	req.ValidateEmail()
	req.ValidatePassword()

	fmt.Printf("%#v", req.Err()) // unexpected EOF; empty email; empty password
}
