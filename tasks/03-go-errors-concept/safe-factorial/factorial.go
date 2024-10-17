package factorial

import (
	"fmt"
)

const maxDepth = 256

var (
	ErrNegativeN error = fmt.Errorf("n should be >= 0")
	ErrTooDeep   error = fmt.Errorf("Too deep factorial. Should de less than %v", maxDepth)
)

// Calculate рекурсивно считает факториал входного числа n.
// Если число меньше нуля, то возвращается ошибка ErrNegativeN.
// Если для вычисления факториала потребуется больше maxDepth фреймов, то Calculate вернёт ErrTooDeep.
func Calculate(n int) (int, error) {
	if n > maxDepth {               
		return 0, ErrTooDeep
	}
	if n < 0 {
		return 0, ErrNegativeN
	}
	res := 1
	i:=1
    for i<n  { 
        res = res * (i+1)
		i = i+1
	}
	return res, nil
}
