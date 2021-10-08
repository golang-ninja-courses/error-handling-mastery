package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var errTarget = stderrors.New("target error")

func checkError(err error) {
	switch {
	case errors.Cause(err) == errTarget:
		fmt.Println("errors.Cause")

	case errors.Is(err, errTarget):
		fmt.Println("errors.Is")

	default:
		fmt.Println("default")
	}
}

func main() {
	{
		err := stderrors.New("target error")
		err = fmt.Errorf("oops: %w", err)
		checkError(err)
	}

	{
		err := fmt.Errorf("oops: %w", errTarget)
		checkError(err)
	}

	{
		err := errors.Wrap(errTarget, "oops")
		checkError(err)
	}
}
