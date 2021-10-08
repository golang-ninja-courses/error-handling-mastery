//nolint:ineffassign,wastedassign
package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func main() {
	if err := collectErrors(); err != nil {
		fmt.Println("errors!", err)
	} else {
		fmt.Println("ok")
	}
}

func collectErrors() error {
	var err error

	if err := foo(); err != nil {
		err = multierror.Append(err, err)
	}

	if err := bar(); err != nil {
		err = multierror.Append(err, err)
	}

	return err
}

func foo() error {
	return errors.New("error from foo")
}

func bar() error {
	return errors.New("error from bar")
}
