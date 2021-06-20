package errorfbench

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.New("invalid token")
	}
}

func BenchmarkErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf("invalid token")
	}
}
