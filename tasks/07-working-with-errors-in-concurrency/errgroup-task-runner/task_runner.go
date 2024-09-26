package tasks

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ErrInvalidWorkersCount = errors.New("invalid workers count")
	ErrFatal               = errors.New("fatal error")
)

type Task interface {
	// Handle выполняет задачу.
	Handle(ctx context.Context) error

	// ExecutionTimeout возвращает промежуток времени,
	// в течение которого задача должна быть выполнена.
	ExecutionTimeout() time.Duration
}

// Run выполняет задачи из очереди tasks с некоторыми условиями:
// – параллельно обрабатываются workersCount задач;
// - если workersCount <= 0, то функция возвращает ErrInvalidWorkersCount;
// – для обработки задачи Task вызывается функция process;
//   - если во время работы process возникла ошибка ErrFatal, то обработка очереди
//     завершается с возвратом этой ошибки;
//   - при любой другой ошибке обработка очереди продолжается.
func Run(ctx context.Context, workersCount int, tasks <-chan Task) error {
	// Для сохранения импортов. Удали эти строки.
	_ = errors.Is
	_ = errgroup.Group{}

	// Реализуй меня.
	return nil
}

// process выполняет задачу task с некоторыми условиями:
// – задача task выполняется не дольше Task.ExecutionTimeout();
// – если во время выполнения задачи возникает ошибка, то она возвращается наружу;
// – при возникновении паники функция возвращает ErrFatal.
func process(ctx context.Context, task Task) (err error) {
	// Реализуй меня.
	return
}
