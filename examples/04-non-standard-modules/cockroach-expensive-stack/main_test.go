package main

import (
	"strconv"
	"testing"
)

// go test -tags pkg -benchmem -bench . > pkg.txt
// go test -tags cockroach -benchmem -bench . > cdb.txt
// benchstat -alpha 1.1 pkg.txt cdb.txt

var depths = []int{1, 2, 4, 8, 16, 32}

func BenchmarkGimmeDeepError(b *testing.B) {
	for _, depth := range depths {
		b.Run(strconv.Itoa(depth), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = GimmeDeepError(depth)
			}
		})
	}
}
