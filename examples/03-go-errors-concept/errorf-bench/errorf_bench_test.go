package errorf_bench

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errors.New("invalid token")
	}
}

func BenchmarkErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Errorf("invalid token")
	}
}
