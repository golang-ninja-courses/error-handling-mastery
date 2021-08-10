package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("error happened")
	fmt.Printf("%+v", err)
}
