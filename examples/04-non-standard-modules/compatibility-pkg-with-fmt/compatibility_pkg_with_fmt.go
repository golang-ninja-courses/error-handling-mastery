package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	stdErr := fmt.Errorf("kek: %w", fmt.Errorf("lol: %w", os.ErrNotExist))
	err := errors.Wrap(errors.Wrap(os.ErrNotExist, "kek"), "lol")
	combined := errors.Wrap(fmt.Errorf("lol: %w", os.ErrNotExist), "kek")
	combined2 := fmt.Errorf("kek: %w", errors.Wrap(os.ErrNotExist, "lol"))

	for _, e := range []error{stdErr, err, combined, combined2} {
		fmt.Printf("Is: %v\n", errors.Is(e, os.ErrNotExist))
	}
}
