package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	file := flag.String("file", "not_found", "file to read")

	flag.Parse()

	f, _ := os.Open(*file)

	reader := bufio.NewScanner(f)

	reader.Split(bufio.ScanWords)

	for reader.Scan() {
		counts[reader.Text()]++
	}

	for x, y := range counts {
		fmt.Printf("item=%s\tcount=%d\n", x, y)
	}
}
