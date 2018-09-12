// Testing some code as I read the chapter
package main

import (
	"fmt"
	"math"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdecafbae)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'b'
	unicode := 'シ'
	newline := '\n'
	fmt.Printf("%d %[1]c %#[1]q\n", ascii)
	fmt.Printf("%d %[1]c %#[1]q\n", unicode)
	fmt.Printf("%d %#[1]q\n", newline)

	fmt.Printf("%e\n", math.MaxFloat32)
	fmt.Printf("%e\n", math.MaxFloat64)
	fmt.Printf("%d\n", math.MinInt32)
	fmt.Printf("%d\n", math.MinInt64)

	var f float32 = 1 << 24
	fmt.Println(f == f+1)

	var g float64 = 1 << 24
	fmt.Println(g == g+1)

}
