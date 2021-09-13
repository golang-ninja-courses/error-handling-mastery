package errors

import "github.com/pkg/errors"

// Wrapf врапит ошибку в текст, при этом прицепляет стектрейс,
// если в цепочке уже нет ошибки со стектрейсом.
func Wrapf(err error, format string, args ...interface{}) error {
	// Для сохранения импортов. Удали эти строки.
	_ = errors.StackTrace{}

	// Реализуй меня.
	return nil
}
