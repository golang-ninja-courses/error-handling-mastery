package errgroup_queue

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	tasksCount   = 100_000
	workersCount = 16

	runTimeout  = 2 * time.Second
	testTimeout = runTimeout + 1*time.Second
)

type task struct {
	executionDuration time.Duration
}

func (t *task) Handle(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(t.executionDuration):
		return nil
	}
}

func (t *task) ExecutionTimeout() time.Duration {
	return 1 * time.Second
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
	tasks := make(chan Task)

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
			errExpected: ErrFatal,
		},
		{
			name: "(workersCount-1) long tasks",
			tasks: queue(func(i int) Task {
				if i < workersCount-1 {
					return &task{executionDuration: 1 * time.Minute}
				}

				return &task{}
			}),
			errExpected: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err         error
				done        = make(chan bool)
				ctx, cancel = context.WithTimeout(context.Background(), runTimeout)
			)
			t.Cleanup(cancel)

			go func() {
				defer func() {
					done <- true
				}()

				err = Run(ctx, workersCount, tt.tasks)
			}()

			select {
			case <-time.After(testTimeout):
				t.Fatal("tested code is too slow")
			case <-done:
			}

			assert.ErrorIs(t, err, tt.errExpected)
		})
	}
}

func TestProcess(t *testing.T) {
	cases := []struct {
		name        string
		task        Task
		errExpected error
	}{
		{
			name:        "normal task",
			task:        &task{},
			errExpected: nil,
		},
		{
			name:        "long task",
			task:        &task{executionDuration: 1 * time.Minute},
			errExpected: nil,
		},
		{
			name:        "panic task",
			task:        &panicTask{},
			errExpected: ErrFatal,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := Process(context.Background(), tt.task)
			assert.ErrorIs(t, err, tt.errExpected)
		})
	}
}

func TestProcess_CanceledCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := Process(ctx, &task{})

	assert.ErrorIs(t, err, context.Canceled)
}
