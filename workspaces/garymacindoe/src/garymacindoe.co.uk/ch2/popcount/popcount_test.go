package popcount_test

import (
    "testing"

    "garymacindoe.co.uk/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(0x1234567890ABCDEF)
    }
}

func BenchmarkPopCountLoop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCountLoop(0x1234567890ABCDEF)
    }
}

func BenchmarkPopCountShift(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCountShift(0x1234567890ABCDEF)
    }
}

func BenchmarkPopCountClear(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCountClear(0x1234567890ABCDEF)
    }
}
