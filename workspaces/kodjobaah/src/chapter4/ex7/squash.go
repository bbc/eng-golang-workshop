package main

import (
	"fmt"
	"unicode"
)

func squash(orig []byte, cur []byte, isSpace boolean, curTotal int, total int) []string {

	if curTotal == total {
		return orig
	}

	if unicode.IsSpace(cur[0]) {
		return inplace(orig, cur[1:], prev, curTotal+1, total)
	}

	if cur[0] == prev {
		return (orig, cur[1:], prev, curTotal+1, total)
	}

	n := make([]string, len(orig)+1, cap(orig)+1)
	copy(n, orig)
	n[len(orig)] = cur[0]
	return inplace(n, cur[1:], cur[0], curTotal+1, total)
}

func main() {
	var test = []byte("This   is the test data")
	fmt.Println(test)
	fmt.Println(squash(test[:1], test[1:], test[0], false, len(test))

}
