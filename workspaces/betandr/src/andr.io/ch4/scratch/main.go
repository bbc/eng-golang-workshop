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

func main() {
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
