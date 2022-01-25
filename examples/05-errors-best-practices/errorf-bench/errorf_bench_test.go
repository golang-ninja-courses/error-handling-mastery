package errorfbench

import (
	"errors"
	"fmt"
	"testing"
)

// ErrGlobal экспортируемая переменная уровня пакета,
// необходимая для предотвращений оптимизаций компилятора.
var ErrGlobal error

func BenchmarkErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrGlobal = errors.New("invalid token")
	}
}

func BenchmarkErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrGlobal = fmt.Errorf("invalid token")
	}
}
