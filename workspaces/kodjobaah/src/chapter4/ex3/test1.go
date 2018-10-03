package main

import (
	"fmt"
)

func modifyValue(s []int) {
	s[1] = 3
	fmt.Printf("In modifyValue: s is %v\n", s)
}

func main() {
	s := []int{1, 2}
	fmt.Printf("In main, before modifyValue: s is %v\n", s)
	modifyValue(s)
	fmt.Printf("In main, after modifyValue: s is %v\n", s)
}
