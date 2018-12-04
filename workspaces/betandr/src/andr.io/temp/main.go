package main

import "fmt"

type T struct{ X int }

func inc1(t *T) {
	t.X++
}

func inc2(t T) T {
	return T{t.X + 1}
}

func main() {
	var a = new(T)
	var b = T{X: 0}

	inc1(a)
	b = inc2(b)

	fmt.Println(a.X)
	fmt.Println(b.X)
}
