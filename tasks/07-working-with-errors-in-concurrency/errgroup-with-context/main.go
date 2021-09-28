package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			fmt.Print("5")
		case <-time.After(3 * time.Second):
			cancel()
		}
	}()

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		select {
		case <-egCtx.Done():
			fmt.Print("1")

		case <-time.After(1 * time.Millisecond):
			return errors.New("2")
		}

		return nil
	})

	eg.Go(func() error {
		select {
		case <-egCtx.Done():
			fmt.Print("3")

		case <-time.After(1 * time.Second):
			fmt.Print("4")
		}

		return nil
	})

	fmt.Print(eg.Wait())
}
