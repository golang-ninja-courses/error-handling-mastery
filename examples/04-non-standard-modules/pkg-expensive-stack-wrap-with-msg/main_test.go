package main

import (
	"strconv"
	"testing"
)

// go test -tags std -benchmem -bench . > std.txt
// go test -tags pkg.msg.stack -benchmem -bench . > pkg-msg-stack.txt
// go test -tags pkg.msg.only -benchmem -bench . > pkg-msg-only.txt
// benchstat -alpha 1.1 std.txt pkg-msg-only.txt
// benchstat -alpha 1.1 std.txt pkg-msg-stack.txt

// ErrGlobal экспортируемая переменная уровня пакета,
// необходимая для предотвращений оптимизаций компилятора.
var ErrGlobal error

var depths = []int{1, 2, 4, 8, 16, 32}

func BenchmarkGimmeDeepError(b *testing.B) {
	for _, depth := range depths {
		b.Run(strconv.Itoa(depth), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ErrGlobal = GimmeDeepError(depth)
			}
		})
	}
}
