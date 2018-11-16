// Write a version of `rotate` that operates in a single pass

package main

import "fmt"

// Rotates a slice of integers `s` at the position `pos` with a single pass
// through the array.
// (NB: This currently doesn't work properly for odd collections)
func rotate(pos int, s []int) {
	k := len(s) - pos
	for i, j := len(s)-1, pos; i > 1; i, j = i-1, j-1 {
		if j <= 0 {
			j = pos
			k = k - pos
		}
		temp := s[i-k]
		s[i-k] = s[i]
		s[i] = temp
	}
}

func main() {
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(4, d[:])
	fmt.Println(d)
}
