package miniword

import (
	"strings"
	"testing"
)

/*
Для авторского решения:

$ go test -benchmem -bench .
BenchmarkDoc-8            163558              6652 ns/op           72608 B/op         18 allocs/op
PASS
ok      05-errors-best-practices/mini-word    1.251s
*/

func BenchmarkDoc(b *testing.B) {
	var w dummyWriter
	text := strings.Repeat("A", 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := NewDocument()

		for i := 2; i <= 3; i++ {
			d.AddPage()
			d.SetActivePage(i)
			for i := 0; i < 100; i++ {
				d.WriteText(text)
			}
		}

		if _, err := d.WriteTo(w); err != nil {
			b.Fatal(err)
		}
	}
}

type dummyWriter struct{}

func (d dummyWriter) Write(_ []byte) (n int, err error) {
	return 0, nil
}
