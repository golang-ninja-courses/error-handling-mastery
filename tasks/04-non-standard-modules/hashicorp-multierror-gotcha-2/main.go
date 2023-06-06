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

	mErr = multierror.Append(mErr, foo())
	mErr = multierror.Append(mErr, bar())

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
