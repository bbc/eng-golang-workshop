// Write a function that counts the number of bits that are different in two
// SHA256 hashes. (see Popcount from Section 2.6.2)

package main

import (
	"fmt"

	"andr.io/ch4/ex4_1/popcount"
)

func main() {
	var i uint64 = 240
	fmt.Printf("%d\n", popcount.Count(i))
}
