//nolint:staticcheck
package main

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func main() {
	if err := collectErrors(); err != nil {
		fmt.Println("errors!")
	} else {
		fmt.Println("ok")
	}
}

func collectErrors() error {
	var err error
	err = multierror.Append(err, foo())
	err = multierror.Append(err, bar())
	return err
}

func foo() error {
	return nil
}

func bar() error {
	return nil
}
