package main

import "fmt"

func notMutate(a [3]int) {
	a[0] = 5
}

func mutate(a [3]int) [3]int {
	a[0] = 5
	return a
}

func mutateByPointer(p *[3]int) {
	p[0] = 8
}

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

// Checks if two slices are equal
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func removeOrdered(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeUnordered(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {

	// remove an item from a slice
	nums1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeOrdered(nums1, 3))
	nums2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeUnordered(nums2, 3))

	s1 := []string{"hello", "world"}
	s2 := []string{"hello", "world"}
	s3 := []string{"bonjour", "le monde"}

	fmt.Println(equal(s1, s2))
	fmt.Println(equal(s1, s3))

	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(3, d[:])
	fmt.Println(d)

	a := [...]int{1, 2, 3}
	fmt.Printf("%T\n", a)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		YEN
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", YEN: "¥"}
	fmt.Printf("%d%s\n", YEN, symbol[YEN])

	r := [...]int{9: -1}
	fmt.Printf("%x\n", r)

	b := [...]int{1, 2, 3}
	notMutate(b)
	fmt.Printf("%x\n", b)

	for i := 0; i > 10; i++ {
		mutate(b) // creates 10 more copies of b!
	}

	c := [...]int{1, 2, 3}
	mutateByPointer(&c)
	fmt.Printf("%x\n", c)
}
