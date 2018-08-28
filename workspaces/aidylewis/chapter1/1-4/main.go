package main

// Usage:
// 	go run main.go 1.txt 2.txt 3.txt

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, fileMap := range counts {
		total := 0
		for _, n := range fileMap {
			total += n
		}

		if total > 1 {
			fmt.Printf("%d %s\n", total, line)

			for file, n := range fileMap {
				fmt.Printf("%d %s\n", n, file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()

		if counts[line] == nil {
			m := make(map[string]int)
			counts[line] = m
		}

		counts[line][f.Name()]++
	}

}
