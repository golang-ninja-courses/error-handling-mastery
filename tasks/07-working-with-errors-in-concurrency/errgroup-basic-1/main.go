package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group

	eg.Go(func() error {
		fmt.Println("1")
		return nil
	})

	eg.Go(func() error {
		fmt.Println("2")
		return nil
	})

	fmt.Println("3")
}
