// Main implementation
package main

import (
	"fmt"

	"andr.io/ch2/ex2_3/popcount"
)

func main() {

	i := uint64(240)
	populationCount := popcount.PopCount(i)
	fmt.Printf("population count of %b is %d\n", i, populationCount)
}
