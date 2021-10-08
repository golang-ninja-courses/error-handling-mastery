package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	workerInterval = 3 * time.Second
	workersCount   = 10
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	// eg, ctx := errgroup.WithContext(ctx)
	// Попробуйте раскомментировать строчку ниже и закомментировать строчку выше,
	// и посмотреть, как поменяется поведение программы.
	var eg errgroup.Group

	for i := 0; i < workersCount; i++ {
		i := i
		eg.Go(func() error {
			t := time.NewTicker(workerInterval)
			defer t.Stop()

			for {
				select {
				case <-ctx.Done():
					return nil
				case <-t.C:
					if err := task(ctx, i); err != nil {
						return fmt.Errorf("do task: %v", err)
					}
				}
			}
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func task(ctx context.Context, taskID int) error {
	fmt.Println(taskID, ": do useful work...")

	if rand.Float64() <= 0.3 { // Возвращаем случайную ошибку в 30% случаях.
		return errors.New("unknown error")
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second):
		return nil
	}
}
