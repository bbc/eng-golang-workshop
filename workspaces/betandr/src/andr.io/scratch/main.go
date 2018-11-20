package main

import "fmt"

func unreachable(i int) int {
	switch i {
	case 0:
		return 1
	case 1:
		return 1
	default:
		return 0
	}

	fmt.Println("not reached")
	return -1
}

func main() {
	var i int

	// always true
	if i != 0 || i != 1 {
		fmt.Println(".")
	}

	// always false
	if i == 0 && i == 1 {
		fmt.Println(".")
	}

	// redundant check
	if i == 0 && i == 0 {
		fmt.Println(".")
	}

	// Self assignment
	j := 42
	j = j

	// overshifting
	var k int8
	k <<= 8

}

// `go vet` works with a single package and does not accept flags
// `go tool vet` more complete, works with files and (recursively) directories and handles options to check by category
