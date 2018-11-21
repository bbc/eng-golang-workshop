// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
)

func squashAdjacentSpaces(phrase []rune) []rune {
	i := 0
	notSpace := true
	for _, c := range phrase {
		if unicode.IsSpace(c) {
			if notSpace {
				phrase[i] = c
				i++
				notSpace = false
			}
		} else {
			phrase[i] = c
			i++
			notSpace = true
		}
	}

	return phrase[:i]
}

func main() {
	phrase := []rune("Finally,  as    the sky      began   to    grow           light")
	unspaced := squashAdjacentSpaces(phrase)
	fmt.Println(string(unspaced))
}
