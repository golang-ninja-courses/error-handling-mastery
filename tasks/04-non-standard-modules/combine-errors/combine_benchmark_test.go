package errors

import (
	"errors"
	"fmt"
	"testing"
)

/*
Для авторского решения:

$ go test -benchmem -bench .
BenchmarkCombine-8       2538996               643.3 ns/op           200 B/op          6 allocs/op
PASS
ok      04-non-standard-modules/combine-errors        2.556s
*/

func BenchmarkCombine(b *testing.B) {
	var (
		err  = errors.New("an error")
		err2 = errors.New("an error 2")
		err3 = errors.New("an error 3")
		err4 = errors.New("an error 4")
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		combinedErr := Combine(err, err2, err3, err4)
		_ = fmt.Sprint(combinedErr)
		_ = fmt.Sprintf("%+v", combinedErr)
	}
}
