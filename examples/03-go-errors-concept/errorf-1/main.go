package main

import (
	"context"
	"errors"
	"fmt"
	"io"
)

func main() {
	err := fmt.Errorf("cannot do operation: %w with %w", io.EOF, context.Canceled)
	fmt.Println(err)                              // cannot do operation: EOF with %!w(*errors.errorString=&{context canceled})
	fmt.Println(errors.Is(err, io.EOF))           // false
	fmt.Println(errors.Is(err, context.Canceled)) // false
}
