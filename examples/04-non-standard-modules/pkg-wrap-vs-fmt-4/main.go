package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var err error // nil error

	{
		err := fmt.Errorf("do operation: %w", err)
		fmt.Println(err, "|", err == nil) // do operation: %!w(<nil>) | false
	}

	{
		err := errors.Wrap(err, "do operation")
		fmt.Println(err, "|", err == nil) // <nil> | true
	}
}
