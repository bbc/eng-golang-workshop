package main

import (
	"fmt"
)

func addValue(s []int) {
	s = append(s, 3)
	fmt.Printf("In addValue: s is %v\n", s)
}

func main() {
	s := []int{1, 2}
	fmt.Printf("In main, before addValue: s is %v\n", s)
	addValue(s)
	fmt.Printf("In main, after addValue: s is %v\n", s)
}
