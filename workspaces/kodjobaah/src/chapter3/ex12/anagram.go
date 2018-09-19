package main

import (
	"fmt"
	"strings"
)

func anagram(first string, second string) bool {

	if len(first) == len(second) {
		for _, v := range first {
			if strings.Contains(second, string(v)) {
				second = strings.Replace(second, string(v), "", -1)
			}

		}
		return len(second) == 0
	} else {
		return false
	}

}

func main() {

	fmt.Printf("Is anagram %v\n", anagram("hey", "yie"))

}
