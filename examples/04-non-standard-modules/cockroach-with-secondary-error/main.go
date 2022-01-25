package main

import (
	"context"
	"fmt"
	"io"

	"github.com/cockroachdb/errors"
)

func main() {
	err := io.EOF
	err = errors.WithSecondaryError(err, context.Canceled)

	fmt.Println(err)
	/*
		EOF
	*/

	fmt.Printf("%+v", err)
	/*
		EOF
		(1) secondary error attachment
		  | context canceled
		  | (1) context canceled
		  | Error types: (1) *errors.errorString
		Wraps: (2) EOF
		Error types: (1) *secondary.withSecondaryError (2) *errors.errorString
	*/

	fmt.Println()
	fmt.Println(errors.Is(err, io.EOF))           // true
	fmt.Println(errors.Is(err, context.Canceled)) // false
}
