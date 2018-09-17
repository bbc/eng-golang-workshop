// Main implementation
package main

import (
	"os"

	"andr.io/ch3/ex3_8/mandelbrot"
)

func main() {
	mandelbrot.Generate128(os.Stdout)
}
