// Нажмите "Отправить" и задача будет принята.
// Авторское решение можно будет посмотреть в решениях.

package errgroup_queue

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

var ErrFatal = errors.New("fatal error")

type Task interface {
	// Handle выполняет задачу.
	Handle(ctx context.Context) error

	// ExecutionTimeout время, в течение которого задача должна быть выполнена.
	ExecutionTimeout() time.Duration
}

// Run выполняет задачи из очереди tasks с некоторыми условиями:
// – параллельно обрабатываются workersCount задач;
// – для обработки задачи Task вызывается функция Process().
func Run(ctx context.Context, workersCount int, tasks <-chan Task) error {
	// Для сохранения импортов. Удали эти строки.
	_ = errors.New
	_ = errgroup.WithContext

	// Реализуй меня.
	return nil
}

// Process выполняет задачу task с некоторыми условиями:
// – задача task выполняется не дольше Task.ExecutionTimeout();
// – если во время выполнения задачи возникает ошибка, то наружу она не возвращается. Считаем, что мы её условно
//	"логируем" и/или "кладём" в какой-нибудь dlq;
// – при возникновении паники возвращает ErrFatal;
// – если контекст ctx закрыт, то возвращает ctx.Err().
func Process(ctx context.Context, task Task) error {
	return nil
}
