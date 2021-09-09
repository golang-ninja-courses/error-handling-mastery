// Нажмите "Отправить" и задача будет принята.
// Авторское решение можно будет посмотреть в решениях.

package errgroup

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

var errFatal error

type Task interface {
	// Handle выполняет задачу.
	Handle(ctx context.Context) error

	// ExecutionTimeout время, в течение которого задача должна быть выполнена.
	ExecutionTimeout() time.Duration
}

// Run выполняет задачи из очереди tasks с некоторыми условиями:
// – параллельно обрабатываются workersCount задач;
// – задача Task выполняется не дольше Task.ExecutionTimeout();
// – если во время выполнения задачи (при вызове метода Task.Handle()) возникает ошибка, то наружу она не возвращается.
//	Считаем, что мы её условно "логируем" и "кладём" задачу в какой-нибудь dlq;
// – при возникновении паники возвращает errFatal;.
func Run(ctx context.Context, workersCount int, tasks <-chan Task) error {
	// Для сохранения импортов. Удали эти строки.
	_ = errors.New
	_ = errgroup.WithContext

	// Реализуй меня.
	return nil
}
