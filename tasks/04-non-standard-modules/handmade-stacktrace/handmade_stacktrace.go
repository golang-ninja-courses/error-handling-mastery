// Подсказка:
// Чтобы правильно решить задачу, рекомендуем немного забежать вперед и посмотреть,
// как сохранение стека реализовано в ошибках из github.com/pkg/errors.
package handmade_stacktrace

const (
	maxStacktraceDepth = 32
)

type Frame uintptr

type StackTrace []Frame

// CallersFrames возвращает стектрейс глубиной не более maxStacktraceDepth.
// Стектрейс, возвращаемый функцией CallerFrames, должен начинаться с того места, где она была вызвана.
func CallersFrames() StackTrace {
	// TODO реализуй меня
	return nil
}