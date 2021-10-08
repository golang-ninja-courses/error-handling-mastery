package errors

import "github.com/pkg/errors"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Wrap оборачивает ошибку в сообщение. Также добавляет к ошибке стектрейс,
// если в цепочке уже нет ошибки со стектрейсом.
func Wrap(err error, msg string) error {
	// Реализуй меня.
	return nil
}
