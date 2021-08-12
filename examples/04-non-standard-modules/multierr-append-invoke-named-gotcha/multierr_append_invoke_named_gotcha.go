package main

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

var errCloserMock = errors.New("close error")

type CloserMock struct{}

func (c *CloserMock) Close() error {
	return errCloserMock
}

func getError() error {
	return nil
}

func good() (err error) {
	err = getError()
	closer := CloserMock{}

	defer multierr.AppendInvoke(&err, multierr.Close(&closer))

	return err
}

func bad() error {
	var err error

	err = getError()
	closer := CloserMock{}

	defer multierr.AppendInvoke(&err, multierr.Close(&closer))

	return err
}

func main() {
	err := good()
	fmt.Printf("good: %v\n", err) // good: close error

	err2 := bad()
	fmt.Printf("bad: %v\n", err2) // bad: <nil>, хотя должен вывести "close error"
}
