package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var ErrNotFound = stderrors.New("not found")

func main() {
	{
		err := fmt.Errorf("index.html: %w", ErrNotFound)
		fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | true
	}

	{
		err := errors.Errorf("index.html: %w", ErrNotFound)
		fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: %!w(*errors.errorString=&{not found}) | false
	}
}
