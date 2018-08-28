package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	s, sep := " ", " "
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("Loop took  %s \n", time.Since(start))

	nextStart := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("Join took %s \n", time.Since(nextStart))

}
