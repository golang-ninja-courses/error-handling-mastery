package main

import "testing"

// ErrGlobal экспортируемая переменная уровня пакета,
// необходимая для предотвращений оптимизаций компилятора.
var ErrGlobal error

func BenchmarkGimmeError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrGlobal = GimmeError()
	}
}

func BenchmarkGimmePkgError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrGlobal = GimmePkgError()
	}
}
