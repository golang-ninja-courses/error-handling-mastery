package factorial

const maxDepth = 256

// Реализуй нас.
var (
	ErrNegativeN = error(nil)
	ErrTooDeep   = error(nil)
)

// Calculate рекурсивно считает факториал входного числа n.
// Если число меньше нуля, то возвращается ошибка ErrNegativeN.
// Если для вычисления факториала потребуется больше maxDepth фреймов, то Calculate вернёт ErrTooDeep.
func Calculate(n int) (int, error) {
	// Реализуй меня.
	return 0, nil
}
