package main

import (
	"errors"
	"fmt"
	"net/http"
	"unsafe"
)

var (
	ErrNotFound  = NewStatusError(http.StatusNotFound, "Not Found")
	ErrNotFound2 = NewStatusError(http.StatusNotFound, "Not Found")
)

type StatusError struct {
	code int
	msg  string
}

func (s *StatusError) Error() string {
	return fmt.Sprintf("%d %s", s.code, s.msg)
}

func NewStatusError(code int, msg string) *StatusError {
	return &StatusError{code: code, msg: msg}
}

func main() {
	fmt.Println(ErrNotFound)                // 404 Not Found
	fmt.Println(unsafe.Sizeof(ErrNotFound)) // 8 (pointer).

	fmt.Println(ErrNotFound == ErrNotFound2)          // false
	fmt.Println(errors.Is(ErrNotFound, ErrNotFound2)) // false

	var se1 error = &StatusError{code: http.StatusNotFound}
	var se2 error = &StatusError{code: http.StatusNotFound}
	fmt.Println(se1 == se2) // false
}
