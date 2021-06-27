package main

import (
	"testing"
)

func BenchmarkGimmeError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GimmeError()
	}
}