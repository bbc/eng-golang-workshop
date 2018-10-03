// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"andr.io/ch2/ex2_4/popcount"
)

// -- Benchmarks --

func BenchmarkPopCountByLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ByLookup(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ByShifting(0x1234567890ABCDEF)
	}
}

// $ go test -cpu=2 -bench=. andr.io/ch2/ex2_4/popcount
// goos: darwin
// goarch: amd64
// pkg: andr.io/ch2/ex2_4/popcount
// BenchmarkPopCountByLookup-2     	100000000	        16.8 ns/op
// BenchmarkPopCountByShifting-2   	20000000	        83.8 ns/op
// PASS
// ok  	andr.io/ch2/ex2_4/popcount	3.463s
