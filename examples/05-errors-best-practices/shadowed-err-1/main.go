package main

import (
	"errors"
	"fmt"
)

var ErrInvalidUserID = errors.New("invalid user id")

type UserID string

type User struct {
	ID UserID
}

//nolint:typecheck
func SaveUser(u User) (err error) {
	if !isValidID(u.ID) {
		err := fmt.Errorf("%w: %v", ErrInvalidUserID, u.ID) // err is shadowed during return
		return
	}

	return saveUser(u)
}

func isValidID(id UserID) bool {
	return false
}

func saveUser(u User) error {
	return nil
}

func main() {
	fmt.Println(SaveUser(User{ID: "XXX"}))
}
