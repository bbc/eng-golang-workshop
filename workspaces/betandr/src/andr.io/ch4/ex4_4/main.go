// Write a version of `rotate` that operates in a single pass

package main

import "fmt"

func rotate(pos int, s []int) {
	for i, j, k := len(s)-1, 3, len(s)-pos; i > 0; i, j = i-1, j-1 {
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
	rotate(3, d[:])
	fmt.Println(d)
}
