package main

import (
	"fmt"
	"io"

	"github.com/cockroachdb/errors"
)

func main() {
	var err error
	err = io.ErrUnexpectedEOF
	err = errors.Mark(err, io.EOF) // Помечаем ошибку как io.EOF
	err = errors.Wrap(err, "something other happened")

	if errors.Is(err, io.EOF) {
		fmt.Println("error is io.EOF") // Выведется
	}

	if errors.Is(err, io.ErrUnexpectedEOF) {
		fmt.Println("error is io.ErrUnexpectedEOF")  // Выведется
	}
}
