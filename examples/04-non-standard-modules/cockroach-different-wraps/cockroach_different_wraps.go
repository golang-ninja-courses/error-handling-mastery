package main

import (
	stdErrors "errors"
	"fmt"

	cockroachErrors "github.com/cockroachdb/errors"
	pkgErrors "github.com/pkg/errors"
)

func main() {
	err := stdErrors.New("an error")

	stdErr := fmt.Errorf("wrap: %w", err)
	fmt.Println(cockroachErrors.UnwrapOnce(stdErr))

	pkgErr := pkgErrors.Wrap(err, "wrap")
	fmt.Println(cockroachErrors.UnwrapOnce(pkgErr))

	cockroachErr := cockroachErrors.Wrap(err, "wrap")
	fmt.Println(cockroachErrors.UnwrapOnce(cockroachErr))
}
