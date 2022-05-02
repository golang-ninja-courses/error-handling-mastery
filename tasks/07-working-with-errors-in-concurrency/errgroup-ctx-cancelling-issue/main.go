package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.TODO()
	action(ctx, "initial action")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		action(ctx, "concurrent action 1")
		return nil
	})
	g.Go(func() error {
		action(ctx, "concurrent action 2")
		return nil
	})

	_ = g.Wait()
	action(ctx, "final action")
}

func action(ctx context.Context, s string) {
	if ctx.Err() != nil {
		fmt.Println(ctx.Err())
		return
	}
	fmt.Println(s)
}
