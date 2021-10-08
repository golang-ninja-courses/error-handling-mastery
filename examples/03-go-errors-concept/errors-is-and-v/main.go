package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	err := io.EOF

	fmt.Println(
		errors.Is(err, io.EOF)) // true
	fmt.Println(
		errors.Is(fmt.Errorf("cannot read file: %v", err), io.EOF)) // false
}
