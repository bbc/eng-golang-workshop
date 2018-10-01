// Package popcount calculates population count of a uint64
//!+
package popcount

// Count returns the population count (number of set bits) of x by clearing.
func Count(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		result++
	}
	return result
}

//!-
