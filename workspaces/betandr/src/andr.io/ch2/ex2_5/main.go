// Main implementation
package main

import (
	"fmt"

	"andr.io/ch2/ex2_5/popcount"
)

func main() {

	i := uint64(240)
	populationCount := popcount.ByLookup(i)
	fmt.Printf("population count of %b by lookup is %d\n", i, populationCount)

	populationCount = popcount.ByClearing(i)
	fmt.Printf("population count of %b by clearing is %d\n", i, populationCount)

	populationCount = popcount.ByShifting(i)
	fmt.Printf("population count of %b by shifting is %d\n", i, populationCount)
}
