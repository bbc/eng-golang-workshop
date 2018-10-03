package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of 1
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcount returns the population count (number of set bits) of x.
func PopCount(x uint8) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))])
}

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%b\n%b\n%t\n%T\n", c1, c2, c1 == c2, c1)
	count := 0
	for i := 0; i < len(c1); i++ {

		count = count + PopCount(c1[i]^c2[i])
		fmt.Printf("c1   =%b\nc2   =%b\nc1^c2=%b\n", c1[i], c2[i], c1[i]^c2[i])
		fmt.Println(PopCount(c1[i] ^ c2[i]))
		fmt.Println("")
	}
	fmt.Println(count)
}
