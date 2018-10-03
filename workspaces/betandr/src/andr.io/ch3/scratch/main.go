// Testing some code as I read the chapter
package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
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

	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[:5], s[7:])
	fmt.Println(s[:])
	fmt.Println("yo" + s[5:])

	fmt.Print("\a")

	y, _ := utf8.DecodeRuneInString("🦄")
	fmt.Println(y)
	fmt.Println(string(y))

	oct31 := 031
	dec25 := 25
	if oct31 == dec25 {
		fmt.Println("Halloween is Christmas")
	}

	fmt.Println("hello\nworld")
	fmt.Println(`hello\nworld`)

	fmt.Println("エリサベス")
	fmt.Println("\u30a8\u30ea\u30b5\u30d9\u30b9")

	s = "hello, エリサベス"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	s = "エリサベス"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("% x\n", r)
	fmt.Println(string(r))

	fmt.Println(strconv.FormatInt(int64(123), 2))
	fmt.Println(fmt.Sprintf("b=%b", int64(123)))

}
