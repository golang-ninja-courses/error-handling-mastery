package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

var (
	errNotFound     = errors.New("not found")
	errUnauthorized = errors.New("unauthorized")
	errUnknown      = errors.New("unknown error")
)

func main() {
	var eg multierror.Group

	eg.Go(func() error { return errNotFound })
	eg.Go(func() error { return errUnauthorized })
	eg.Go(func() error { return errUnknown })

	err := eg.Wait()
	fmt.Println(errors.Is(err, errNotFound))     // true
	fmt.Println(errors.Is(err, errUnauthorized)) // true
	fmt.Println(errors.Is(err, errUnknown))      // true
}
