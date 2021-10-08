package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	errNotFound     = errors.New("not found")
	errUnauthorized = errors.New("unauthorized")
	errUnknown      = errors.New("unknown error")
)

func main() {
	var eg errgroup.Group

	eg.Go(func() error { return errNotFound })
	eg.Go(func() error { return errUnauthorized })
	eg.Go(func() error { return errUnknown })

	switch err := eg.Wait(); {
	case errors.Is(err, errNotFound):
		fmt.Println("1")

	case errors.Is(err, errUnauthorized):
		fmt.Println("2")

	case errors.Is(err, errUnknown):
		fmt.Println("3")
	}
}
