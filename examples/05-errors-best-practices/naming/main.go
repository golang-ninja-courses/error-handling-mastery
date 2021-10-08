// nolint:deadcode,unused,varcheck
package naming

import (
	"errors"
	"fmt"
)

var (
	// Sentinel errors.
	// Название начинается с приставки err или Err.
	errNotFound = errors.New("error not found") // Неэкспортируемая.
	ErrNotFound = errors.New("error not found") // Экспортируемая.
)

// Кастомный тип ошибки. Название заканчивается на Error.
type NotFoundError struct {
	page string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("page %q not found", e.page)
}

var (
	// Sentinel errors.
	// Точно так же название начинается с приставки err или Err.
	errNotFound2 = &NotFoundError{"https://www.golang-courses.ru/"} // Неэкспортируемая.
	ErrNotFound2 = &NotFoundError{"https://www.golang-courses.ru/"} // Экспортируемая.
)
