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
	var mErr *multierror.Error

	if err := foo(); err != nil {
		mErr = multierror.Append(mErr, err)
	}

	if err := bar(); err != nil {
		mErr = multierror.Append(mErr, err)
	}

	if mErr != nil {
		mErr.ErrorFormat = func(errors []error) string {
			return fmt.Sprintf("%v", errors)
		}
	}

	return mErr
}

func foo() error {
	return nil
}

func bar() error {
	return nil
}
