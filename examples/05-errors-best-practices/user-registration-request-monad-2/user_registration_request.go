package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
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
	if u.err != nil {
		return
	}
	u.err = json.NewDecoder(r).Decode(u)
}

func (u *UserRegistrationRequest) ValidateEmail() {
	if u.err != nil {
		return
	}

	if u.Email == "" {
		u.err = errors.New("empty email")
	}
}

func (u *UserRegistrationRequest) ValidatePassword() {
	if u.err != nil {
		return
	}

	if u.Password == "" {
		u.err = errors.New("empty password")
	}
}

func main() {
	var req UserRegistrationRequest
	req.Unmarshal(strings.NewReader(`{"email":"bob@gmail.com","password":""}`))
	req.ValidateEmail()
	req.ValidatePassword()

	fmt.Println(req.Err()) // empty password
}
