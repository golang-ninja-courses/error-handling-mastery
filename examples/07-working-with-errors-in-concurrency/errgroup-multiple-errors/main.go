package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group

	eg.Go(func() error {
		return errors.New("first error")
	})

	eg.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("second error")
	})

	eg.Go(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("third error")
	})

	fmt.Println(eg.Wait()) // first error
}
