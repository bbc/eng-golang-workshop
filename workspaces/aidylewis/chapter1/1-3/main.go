package main

// Usage:
// 	go run main.go {1..19300}

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func inefficient() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

func moreInefficient() {
	strings.Join(os.Args[1:], " ")
}

func main() {
	startIneff := time.Now()
	inefficient()
	secsIneff := time.Since(startIneff).Seconds()
	startEff := time.Now()
	moreInefficient()
	secsEff := time.Since(startEff).Seconds()
	fmt.Println(secsIneff)
	fmt.Println(secsEff)
}
