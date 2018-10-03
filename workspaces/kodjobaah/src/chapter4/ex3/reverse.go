// reverse reverses a slice of ints in place
package main

import "fmt"

func reverse(s []int) {
	fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions
	//	p := s[:3]
	//	fmt.Println(p)
	//	reverse(p)
	//	fmt.Print("---")
	//	fmt.Print(s)
	//	fmt.Println("")
	reverse(s[3:])
	fmt.Print("---")
	fmt.Print(s)
	fmt.Println("")
	reverse(s)

}
