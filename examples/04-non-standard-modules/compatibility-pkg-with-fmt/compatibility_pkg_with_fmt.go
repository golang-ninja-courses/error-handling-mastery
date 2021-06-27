package main

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func main() {
	stdErr := fmt.Errorf("kek: %w", fmt.Errorf("lol: %w", io.EOF))
	err := errors.Wrap(errors.Wrap(io.EOF, "kek"), "lol")
	combined := errors.Wrap(fmt.Errorf("lol: %w", io.EOF), "kek")
	combined2 := fmt.Errorf("kek: %w", errors.Wrap(io.EOF, "lol"))

	for _, e := range []error{stdErr, err, combined, combined2} {
		fmt.Printf("Is: %v\n", errors.Is(e, io.EOF))
	}
}
