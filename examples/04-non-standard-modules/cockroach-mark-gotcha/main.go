package main

import (
	stderrors "errors"
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

	err := io.ErrUnexpectedEOF
	err = errors.Mark(err, err1)
	err = errors.Wrap(err, "something other happened")

	if errors.Is(err, err1) { // Ожидаемо true.
		fmt.Println("err is err1")
	}

	if errors.Is(err, err2) { // Не очень ожидаемо true.
		fmt.Println("err is err2")
	}

	if errors.Is(err1, err2) { // Ещё менее ожидаемо true.
		fmt.Println("err1 is err2")
	}

	if stderrors.Is(err1, err2) { // Ожидаемо false.
		fmt.Println("err1 is err2")
	}
}
