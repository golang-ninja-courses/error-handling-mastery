package monad

import "errors"

var ErrNoMonadValue = errors.New("no monad value")

// M представляет собой монаду.
type M struct {
	err error
	v   any
}

// Bind применяет функцию f к значению M, возвращая новую монаду.
// Если M невалидна, то Bind эффекта не имеет.
func (m M) Bind(f func(v any) M) M {
	// Реализуй меня.
	return M{}
}

// Unpack возвращает значение и ошибку, хранимые в монаде.
func (m M) Unpack() (any, error) {
	// Реализуй меня.
	return nil, nil
}

// Unit конструирует M на основе значения v.
func Unit(v any) M {
	// Реализуй меня.
	return M{}
}

// Err конструирует "невалидную" монаду M.
func Err(err error) M {
	// Реализуй меня.
	return M{}
}
