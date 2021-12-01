package tasks

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

const defaultRunTimeout = 5 * time.Second

func TestRun(t *testing.T) {
	var c counter

	inc := func() error {
		c.Inc()
		return nil
	}

	doNothing := func() error {
		return nil
	}

	withPanic := func() error {
		panic("panic!")
	}

	withError := func() error {
		_ = inc()
		return errors.New("connection refused")
	}

	cases := []struct {
		name             string
		workers          int
		tasks            []Task
		runTimeout       time.Duration
		expectedErr      error
		expectedExecTime time.Duration // Всегда берём с запасом из-за погрешности на среду выполнения.
		checkCounter     bool
	}{
		{
			name:        "invalid workers count",
			workers:     -1,
			tasks:       newTasks(1000, doNothing, time.Second/2, time.Second),
			expectedErr: ErrInvalidWorkersCount,
		},
		{
			name:        "no workers",
			workers:     0,
			tasks:       newTasks(1000, doNothing, time.Second/2, time.Second),
			expectedErr: ErrInvalidWorkersCount,
		},
		{
			name:             "1 worker X 60 tasks (50ms) = 3s",
			workers:          1,
			tasks:            newTasks(60, inc, 50*time.Millisecond, time.Second),
			expectedErr:      nil,
			expectedExecTime: 4 * time.Second,
			checkCounter:     true,
		},
		{
			name:             "2 workers X 60 tasks (50ms) = 1.5s",
			workers:          2,
			tasks:            newTasks(60, inc, 50*time.Millisecond, time.Second),
			expectedErr:      nil,
			expectedExecTime: 2 * time.Second,
			checkCounter:     true,
		},
		{
			name:             "60 workers X 600 tasks (50ms) = 500ms",
			workers:          60,
			tasks:            newTasks(600, inc, 50*time.Millisecond, time.Second),
			expectedErr:      nil,
			expectedExecTime: time.Second,
			checkCounter:     true,
		},
		{
			name:             "16 workers, 200 long tasks, cancellation by root ctx",
			workers:          16,
			tasks:            newTasks(200, doNothing, time.Minute, 3*time.Minute),
			runTimeout:       time.Second,
			expectedErr:      context.DeadlineExceeded,
			expectedExecTime: 2 * time.Second,
		},
		{
			name:             "16 workers, 16 long tasks, cancellation by self timeout",
			workers:          16,
			tasks:            newTasks(16, doNothing, time.Minute, time.Second),
			expectedErr:      nil,
			expectedExecTime: 2 * time.Second,
		},
		{
			name:    "16 workers, 15 long tasks & 1 task with panic",
			workers: 16,
			tasks: append(
				newTasks(15, doNothing, time.Minute, 3*time.Minute),
				newTask(withPanic, time.Second, 3*time.Second),
			),
			expectedErr:      ErrFatal,
			expectedExecTime: 2 * time.Second,
		},
		{
			name:             "16 workers, 16 errored tasks, no error from Run",
			workers:          16,
			tasks:            newTasks(15, withError, 50*time.Millisecond, time.Second),
			expectedErr:      nil,
			expectedExecTime: time.Second / 2,
			checkCounter:     true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer goleak.VerifyNone(t)
			defer c.Reset()

			timeout := tt.runTimeout
			if timeout == 0 {
				timeout = defaultRunTimeout
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			errc := make(chan error, 1)
			go func() {
				errc <- Run(ctx, tt.workers, queue(ctx, tt.tasks...))
			}()

			start := time.Now()

			select {
			case <-time.After(defaultRunTimeout):
				t.Fatal("tested code is too slow: Run blocked?")

			case err := <-errc:
				require.ErrorIs(t, err, tt.expectedErr)

				if tt.expectedExecTime != 0 {
					assert.LessOrEqual(t, time.Since(start), tt.expectedExecTime,
						"workers did not work in parallel")
				}
				if tt.checkCounter {
					assert.Equal(t, len(tt.tasks), c.Value(),
						"not all tasks were processed")
				}
			}
		})
	}
}

type counter struct {
	i  int
	mu sync.RWMutex
}

func (c *counter) Inc() {
	c.mu.Lock()
	c.i++
	c.mu.Unlock()
}

func (c *counter) Reset() {
	c.mu.Lock()
	c.i = 0
	c.mu.Unlock()
}

func (c *counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.i
}

func newTasks(n int, f func() error, duration, timeout time.Duration) []Task {
	tasks := make([]Task, n)
	for i := 0; i < n; i++ {
		tasks[i] = newTask(f, duration, timeout)
	}
	return tasks
}

type task struct {
	f        func() error
	timeout  time.Duration
	duration time.Duration
}

func newTask(f func() error, d, t time.Duration) task {
	return task{f: f, duration: d, timeout: t}
}

func (t task) ExecutionTimeout() time.Duration {
	return t.timeout
}

func (t task) Handle(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(t.duration):
		return t.f()
	}
}

func queue(ctx context.Context, tasks ...Task) <-chan Task {
	tasksCh := make(chan Task)

	go func() {
		defer close(tasksCh)

		for _, t := range tasks {
			select {
			case <-ctx.Done():
				return
			case tasksCh <- t:
			}
		}
	}()

	return tasksCh
}
