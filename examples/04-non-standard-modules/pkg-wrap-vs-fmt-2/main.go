package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var ErrNotFound = stderrors.New("not found")

func main() {
	{
		err := fmt.Errorf("index.html: %v", ErrNotFound)
		fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | false
	}

	{
		err := errors.Errorf("index.html: %v", ErrNotFound)
		fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | false
	}
}
