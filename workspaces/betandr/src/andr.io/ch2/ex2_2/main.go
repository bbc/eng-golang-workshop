// Main implementation
package main

import (
	"fmt"
	"os"
	"strconv"

	t "andr.io/ch2/ex2_1"
	w "andr.io/ch2/ex2_2/weightconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		n, _ := strconv.Atoi(arg)
		fmt.Printf("-= converting %d =- \n", n)

		c := t.Celsius(n)
		fmt.Printf("%s is %s\n", c, t.CToF(c))

		f := t.Fahrenheit(n)
		fmt.Printf("%s is %s\n", f, t.FToC(f))

		k := t.Kelvin(n)
		fmt.Printf("%s is %s\n", k, t.KToF(k))

		g := w.Kilo(n)
		fmt.Printf("%s is %s\n", g, w.KToP(g))

		l := w.Pound(n)
		fmt.Printf("%s is %s\n", l, w.PToS(l))

		s := w.Stone(n)
		fmt.Printf("%s is %s\n", s, w.SToP(s))
	}

}
