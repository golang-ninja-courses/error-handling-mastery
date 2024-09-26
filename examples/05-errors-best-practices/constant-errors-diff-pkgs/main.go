//nolint:gocritic,ineffassign,staticcheck,wastedassign
package main

import (
	"fmt"

	"github.com/golang-ninja-courses/error-handling-mastery/examples/05-errors-best-practices/constant-errors-diff-pkgs/pkga"
	"github.com/golang-ninja-courses/error-handling-mastery/examples/05-errors-best-practices/constant-errors-diff-pkgs/pkgb"
)

func main() {
	fmt.Println(pkga.ErrUnknownData == pkgb.ErrUnknownData) // true - created by global type

	// invalid operation: pkga.ErrInvalidHost == pkgb.ErrInvalidHost (mismatched types pkga.err and pkgb.err)
	// fmt.Println(pkga.ErrInvalidHost == pkgb.ErrInvalidHost)

	fmt.Println(error(pkga.ErrInvalidHost) == error(pkga.ErrInvalidHost)) // true - one type
	fmt.Println(error(pkga.ErrInvalidHost) == error(pkgb.ErrInvalidHost)) // false - diff types

	if err := foo(); err != nil {
		fmt.Println(err) // invalid host
	}

	// cannot use foo() (type error) as type pkga.err in assignment: need type assertion
	// err := pkga.ErrInvalidHost
	// err = foo()

	var err error = pkga.ErrInvalidHost
	err = foo()
	_ = err
}

func foo() error {
	return pkga.ErrInvalidHost
}
