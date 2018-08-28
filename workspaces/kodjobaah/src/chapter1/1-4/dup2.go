// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
			fmt.Printf("%s\n", line)
		for c, _ := range n {
			fmt.Printf("\t%s\n", c)
		}
	}
	
}

func countLines(f *os.File, counts map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		fmap, state := counts[txt]
		if ( state == false) {
			fmap = make(map[string]bool)
		}
		fmap[f.Name()] = true
		counts[txt] = fmap
	}
	// NOTE: ignoring potential errors from input.Err()
}
