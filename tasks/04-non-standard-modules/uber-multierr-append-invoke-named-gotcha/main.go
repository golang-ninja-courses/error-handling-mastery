package main

import (
	"fmt"
	"io"

	"go.uber.org/multierr"
)

func main() {
	err := good()
	fmt.Printf("good: %v\n", err) // good: unexpected EOF

	err2 := bad()
	fmt.Printf("bad: %v\n", err2) // bad: <nil>, хотя должен вывести "unexpected EOF"
}

var errInvoker = multierr.Invoke(func() error { return io.ErrUnexpectedEOF })

func good() (err error) {
	err = getError()
	defer multierr.AppendInvoke(&err, errInvoker)
	return nil
}

// https://golang.org/ref/spec#Return_statements
func bad() error {
	err := getError()
	defer multierr.AppendInvoke(&err, errInvoker)
	return err
}

func getError() error {
	return nil
}
