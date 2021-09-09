package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var errInternal = stderrors.New("ooops, an error on level 1")

func GimmeDeepError(depth int) error {
	if depth == 1 {
		return errors.WithStack(errInternal)
	}
	return errors.WithMessagef(GimmeDeepError(depth-1), "error happened on level %d", depth)
}

func main() {
	fmt.Printf("%+v", GimmeDeepError(5))
}
