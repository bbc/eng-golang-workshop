// Write a version of `rotate` that operates in a single pass

package main

import "fmt"

// Rotates a slice of integers `s` at the position `pos` with a single pass
// through the array, swapping from the last position to the first position
// until all items are in order, running in O(n) time and using O(n) space.
func rotate(pos int, s []int) {
	k := len(s) - pos
	stop := 0
	if len(s)%2 == 0 {
		stop = 1
	}

	for i, j := len(s)-1, pos; i > stop; i, j = i-1, j-1 {
		if j <= 0 {
			j = pos
			k = k - pos
		}
		l := i - k
		if i <= 2 {
			l = 0
		}
		temp := s[l]
		s[l] = s[i]
		s[i] = temp
		fmt.Printf("swap %d with %d\n", l, i)
	}
}

func main() {
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(4, d[:])
	fmt.Println(d)
}
