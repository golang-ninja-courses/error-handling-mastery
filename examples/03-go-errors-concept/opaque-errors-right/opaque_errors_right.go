package main

import (
	"errors"
	"fmt"
)

type isTemporary interface {
	IsTemporary() bool
}

type MyError struct {
	cause error
}

func (m *MyError) Error() string {
	return fmt.Sprintf("my error: %s", m.cause.Error())
}

func (m *MyError) Unwrap() error {
	return m.cause
}

func (m *MyError) Cause() error {
	return m.cause
}

func (m *MyError) IsTemporary() bool {
	return true
}

func IsTemporaryWrong(err error) bool {
	// Приводим ошибку к интерфейсу, тем самым проверяя поведение.
	e, ok := err.(isTemporary)
	return ok && e.IsTemporary()
}

func UnwrapOnce(err error) error {
	switch e := err.(type) {
	case interface{ Cause() error }:
		return e.Cause()
	case interface{ Unwrap() error }:
		return e.Unwrap()
	}
	return nil
}

func IsTemporary(err error) bool {
	for c := err; c != nil; c = UnwrapOnce(c) {
		e, ok := c.(isTemporary)
		if ok && e.IsTemporary() {
			return true
		}
	}

	return false
}

func main() {
	err := errors.New("an error")
	errMy := &MyError{cause: err}
	errWrapped := fmt.Errorf("wrapped error: %w", errMy)

	fmt.Println("wrong:")
	fmt.Printf("%v\n", IsTemporaryWrong(err))
	fmt.Printf("%v\n", IsTemporaryWrong(errMy))
	fmt.Printf("%v\n", IsTemporaryWrong(errWrapped))

	fmt.Println("\nright:")
	fmt.Printf("%v\n", IsTemporary(err))
	fmt.Printf("%v\n", IsTemporary(errMy))
	fmt.Printf("%v\n", IsTemporary(errWrapped))
}
