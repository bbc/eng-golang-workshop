package ex3_5

import (
	"chapter2/ex3_5/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(9223372036854775807)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {

	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(9223372036854775807)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(9223372036854775807)
	}
}

func BenchmarkPopCountClearRightmost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearRightmost(9223372036854775807)
	}
}
