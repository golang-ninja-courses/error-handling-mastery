package main

import (
	"strconv"
	"testing"
)

func BenchmarkGimmeDeepError(b *testing.B) {
	depths := []int{1, 2, 4, 8, 16, 32}
	for _, depth := range depths {
		b.Run(strconv.Itoa(depth), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = GimmeDeepError(depth)
			}
		})
	}
}