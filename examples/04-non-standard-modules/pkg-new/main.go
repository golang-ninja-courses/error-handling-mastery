package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("error happened")
	fmt.Printf("%+v", err)

	sErr := stderrors.New("error happened")
	fmt.Printf("\n---\n%+v", sErr)
}
