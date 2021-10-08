package main

import (
	stderrors "errors"
	"fmt"

	cdberrors "github.com/cockroachdb/errors"
	pkgerrors "github.com/pkg/errors"
)

func main() {
	err := stderrors.New("an error")

	stdErr := fmt.Errorf("wrap: %w", err)
	fmt.Println(cdberrors.UnwrapOnce(stdErr))

	pkgErr := pkgerrors.Wrap(err, "wrap 1")
	fmt.Println(cdberrors.UnwrapOnce(pkgErr))

	cockroachErr := cdberrors.Wrap(err, "wrap 2")
	fmt.Println(cdberrors.UnwrapOnce(cockroachErr))
}
