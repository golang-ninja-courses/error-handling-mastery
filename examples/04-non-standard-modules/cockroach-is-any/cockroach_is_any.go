package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cockroachdb/errors"
)

func main() {
	errs := []error{context.Canceled, io.EOF, os.ErrClosed, os.ErrNotExist, errors.New("unknown error")}

	for _, err := range errs {
		if errors.IsAny(err, errs...) {
			fmt.Println("gotcha!")
		}
	}
}
