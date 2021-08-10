package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type M = UserRegistrationRequestMonad

type UserRegistrationRequestMonad struct {
	req UserRegistrationRequest
	err error
}

func (u M) Unpack() (UserRegistrationRequest, error) {
	return u.req, u.err
}

func (u M) Bind(f func(req UserRegistrationRequest) M) M {
	if u.err != nil {
		return u
	}
	return f(u.req)
}

func unmarshalRequest(r io.Reader) func(req UserRegistrationRequest) M {
	return func(req UserRegistrationRequest) M {
		if err := json.NewDecoder(r).Decode(&req); err != nil {
			return M{err: err}
		}
		return M{req: req}
	}
}

func validateEmail(req UserRegistrationRequest) M {
	if req.Email == "" {
		return M{err: errors.New("empty email")}
	}
	return M{req: req}
}

func validatePassword(req UserRegistrationRequest) M {
	if req.Password == "" {
		return M{err: errors.New("empty password")}
	}
	return M{req: req}
}

func main() {
	req, err := UserRegistrationRequestMonad{}.
		Bind(unmarshalRequest(strings.NewReader(`{"email":"bob@gmail.com","password":"bob"}`))).
		Bind(validateEmail).
		Bind(validatePassword).
		Unpack()

	fmt.Println(err)       // nil
	fmt.Printf("%#v", req) // main.UserRegistrationRequest{Email:"bob@gmail.com", Password:"bob"}
}
