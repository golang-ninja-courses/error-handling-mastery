package main

import (
	stdErrors "errors"
	"fmt"
	"io"

	"github.com/cockroachdb/errors"
)

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func main() {
	err1 := &NotFoundError{"object not found"}
	err2 := &NotFoundError{"object not found"}

	var err error
	err = io.ErrUnexpectedEOF
	err = errors.Mark(err, err1)
	err = errors.Wrap(err, "something other happened")

	if errors.Is(err, err1) {
		fmt.Println("err is err1") // Ожидаемо выведется
	}

	if errors.Is(err, err2) {
		fmt.Println("err is err2") // <--- Не очень ожидаемо выведется
	}

	if errors.Is(err1, err2) {
		fmt.Println("err1 is err2") // <--- Ещё менее ожидаемо выведется
	}

	if stdErrors.Is(err1, err2) {
		fmt.Println("err1 is err2") // Ожидаемо не выведется
	}
}
