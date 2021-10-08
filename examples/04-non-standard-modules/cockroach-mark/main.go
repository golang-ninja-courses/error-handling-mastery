package main

import (
	"fmt"
	"io"

	"github.com/cockroachdb/errors"
)

func main() {
	err := io.ErrUnexpectedEOF
	err = errors.Mark(err, io.EOF) // Помечаем ошибку как io.EOF.
	err = errors.Wrap(err, "something other happened")

	if errors.Is(err, io.EOF) { // true
		fmt.Println("error is io.EOF")
	}

	if errors.Is(err, io.ErrUnexpectedEOF) { // true
		fmt.Println("error is io.ErrUnexpectedEOF")
	}
}
