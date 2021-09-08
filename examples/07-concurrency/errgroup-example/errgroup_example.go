package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	defaultQueueWorkersCount  = 4
	scheduleQueueWorkersCount = 3
	paymentsQueueWorkersCount = 2
)

var (
	errTerminated = errors.New("process was terminated")
)

type queue struct{}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Run(ctx context.Context, tickPeriod time.Duration) error {
	errChan := make(chan error, 1)
	ticker := time.NewTicker(tickPeriod)
	defer ticker.Stop()

	for range ticker.C {
		go func() {
			errChan <- q.process(ctx)
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (q *queue) process(ctx context.Context) error {
	// Как-то обрабатываем очередь.
	return nil
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < defaultQueueWorkersCount; i++ {
		defaultQueue := NewQueue()
		g.Go(func() error {
			return defaultQueue.Run(ctx, 3*time.Second)
		})
	}

	for i := 0; i < scheduleQueueWorkersCount; i++ {
		scheduleQueue := NewQueue()
		g.Go(func() error {
			return scheduleQueue.Run(ctx, 3*time.Second)
		})
	}

	for i := 0; i < paymentsQueueWorkersCount; i++ {
		paymentsQueue := NewQueue()
		g.Go(func() error {
			return paymentsQueue.Run(ctx, 3*time.Second)
		})
	}

	g.Go(func() error {
		signalChan := make(chan os.Signal, 1)

		signal.Notify(signalChan, syscall.SIGINT, os.Interrupt, syscall.SIGTERM)

		select {
		case s := <-signalChan:
			fmt.Printf("signal caught: %v\n", s)
			return errTerminated
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, errTerminated) && !errors.Is(err, context.Canceled) {
		fmt.Printf("errgroup returned error: %+v", err)
		return
	}

	fmt.Println("queue handlers were gracefully shut down")
}
