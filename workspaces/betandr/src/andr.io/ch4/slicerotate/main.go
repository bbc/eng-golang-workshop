package main

import "fmt"

// Reverse a slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Rotate a slice at a given position
func rotate(pos int, s []int) {
	reverse(s[:pos])
	reverse(s[pos:])
	reverse(s)
}

func main() {
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(3, d[:])
	fmt.Println(d)
}
