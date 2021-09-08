package errgroup

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	tasksCount    = 250_000
	queueCapacity = 100
	workersCount  = 16

	runTimeout  = 2 * time.Second
	testTimeout = runTimeout + 1*time.Second
)

type task struct {
	executionDuration time.Duration
}

func (t *task) Handle(_ context.Context) error {
	time.Sleep(t.ExecutionTimeout())
	return nil
}

func (t *task) ExecutionTimeout() time.Duration {
	return t.executionDuration
}

type errorTask struct {
	task
}

func (t *errorTask) Handle(_ context.Context) error {
	return errors.New("errorTask error")
}

type panicTask struct {
	task
}

func (t *panicTask) Handle(_ context.Context) error {
	panic("panic!")
}

func queue(generateTask func(i int) Task) <-chan Task {
	tasks := make(chan Task, queueCapacity)

	go func() {
		for i := 0; i < tasksCount; i++ {
			tasks <- generateTask(i)
		}

		close(tasks)
	}()

	return tasks
}

func TestRun(t *testing.T) {
	cases := []struct {
		name        string
		tasks       <-chan Task
		errExpected error
	}{
		{
			name: "normal tasks",
			tasks: queue(func(_ int) Task {
				return &task{}
			}),
			errExpected: nil,
		},
		{
			name: "error tasks",
			tasks: queue(func(_ int) Task {
				return &errorTask{}
			}),
			errExpected: nil,
		},
		{
			name: "long tasks",
			tasks: queue(func(_ int) Task {
				return &task{executionDuration: 1 * time.Minute}
			}),
			errExpected: context.DeadlineExceeded,
		},
		{
			name: "panic tasks",
			tasks: queue(func(_ int) Task {
				return &panicTask{}
			}),
			errExpected: errFatal,
		},
		{
			name: "mixed",
			tasks: queue(func(i int) Task {
				divider := tasksCount / (workersCount - 1)

				switch i % divider {
				case divider - 1:
					return &task{executionDuration: 1 * time.Minute}
				default:
					return &task{}
				}
			}),
			errExpected: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan bool)
			ctx, cancel := context.WithTimeout(context.Background(), runTimeout)
			t.Cleanup(cancel)

			go func() {
				defer func() {
					done <- true
				}()

				err := Run(ctx, workersCount, tt.tasks)

				assert.ErrorIs(t, err, tt.errExpected)
			}()

			select {
			case <-time.After(testTimeout):
				t.Fatal("tested code is too slow")
			case <-done:
			}
		})
	}
}
