//General Purpose unit conversion application
package main

import (
	"chapter2/ex2/conversion"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]
	if len(files) == 0 {
		conversion.ConvertUnits(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unitConversion: %v\n", err)
				continue
			}
			conversion.ConvertUnits(f)
			f.Close()
		}

	}
}
