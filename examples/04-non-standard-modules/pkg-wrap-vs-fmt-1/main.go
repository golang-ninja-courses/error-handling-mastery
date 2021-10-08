package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var ErrNotFound = stderrors.New("not found")

func main() {
	{
		err := fmt.Errorf("%w: index.html", ErrNotFound)
		fmt.Println(err) // not found: index.html

		err = fmt.Errorf("in the middle: %w: index.html", ErrNotFound)
		fmt.Println(err) // in the middle: not found: index.html

		err = fmt.Errorf("index.html: %w", ErrNotFound)
		fmt.Println(err) // index.html: not found
	}

	{
		err := errors.Wrap(ErrNotFound, "index.html")
		fmt.Println(err) // index.html: not found
	}
}
