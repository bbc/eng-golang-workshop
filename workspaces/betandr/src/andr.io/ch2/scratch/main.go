// Testing some code as I read the chapter
package main

import (
	"fmt"
	"strconv"
)

var i = 0

func main() {
	i, j := 1, 1
	fmt.Println(strconv.Itoa(i) + " - " + strconv.Itoa(j))
	p := &i
	fmt.Println(*p)
	q := f()
	r := f()
	fmt.Println(q == r)

	s := 1
	incr(&s)
	fmt.Println(incr(&s))

	var t = new(int)
	*t = 4
	fmt.Println(*t)
	// ...is like:
	var u int
	u = 5
	fmt.Println(u)

	v := 6
	v++
	v--
	fmt.Println(v)

	x := gcd(42, 7)
	fmt.Println(x)

	y := fib(6)
	fmt.Println(y)

	// implicit:
	nums := []int{9}
	fmt.Println(nums[0])

	// explicit:
	nums[0] = 10
	fmt.Println(nums[0])

	/*
		comparibility needs assignability, where `nums[0] = 11` would be valid but
		`nums[0]` == "eleven" would not be
	*/
	fmt.Println(nums[0] == 10)
}

//
func f() *int {
	v := 2
	return &v
}

// Increment the value referenced by p and return reference
func incr(p *int) int {
	*p++
	return *p
}

// Compute the greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Compute the nth fibbonaci number
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
