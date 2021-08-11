package main

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
)

func getError() error {
	return errors.New("an error")
}

func good() (err error) {
	defer multierr.AppendInvoke(&err, multierr.Invoke(getError))
	return err
}

func bad() (err error) {
	defer multierr.AppendInto(&err, getError())
	return err
}

func main() {
	err := good()
	fmt.Printf("good: %v\n", err)

	fmt.Println()

	err2 := bad()
	fmt.Printf("bad: %v\n", err2)
}
