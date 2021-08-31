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

func good() (err error) {
	c := new(CloserMock)
	defer multierr.AppendInvoke(&err, multierr.Close(c))

	c.SetErr(io.ErrUnexpectedEOF)
	return nil
}

func bad() (err error) {
	c := new(CloserMock)
	defer multierr.AppendInto(&err, c.Close())

	c.SetErr(io.ErrUnexpectedEOF)
	return nil
}
