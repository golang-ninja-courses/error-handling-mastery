package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	errProgramInterrupted   = errors.New("program interrupted")
	errOrchestrationTimeout = errors.New("orchestration timeout")
	errWorkTimeout          = errors.New("work timeout")
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel(errProgramInterrupted)
	}()

	err := orchestrate(ctx)
	fmt.Println(err)
	// aborting work: context canceled, because of program interrupted
	// or
	// aborting work: context deadline exceeded, because of orchestration timeout
}

func orchestrate(ctx context.Context) error {
	ctx, cancel := context.WithTimeoutCause(ctx, 3*time.Second, errOrchestrationTimeout)
	defer cancel()
	return doWork(ctx)
}

func doWork(ctx context.Context) error {
	ctx, cancel := context.WithTimeoutCause(ctx, 10*time.Second, errWorkTimeout)
	defer cancel()

	<-ctx.Done()
	return fmt.Errorf("aborting work: %v, because of %w", ctx.Err(), context.Cause(ctx))
}
