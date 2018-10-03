package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

//If contains multiple spaces in front does not squash
func squash(orig []byte, cur []byte, isSpace bool, curTotal int, total int) []byte {

	if curTotal == (total - 1) {
		return orig
	}

	var b = []byte{cur[0]}
	var r, _ = utf8.DecodeRune(b)

	if unicode.IsSpace(r) && isSpace {
		return squash(orig, cur[1:], true, curTotal+1, total)
	}

	n := make([]byte, len(orig)+1, cap(orig)+1)
	copy(n, orig)
	n[len(orig)] = cur[0]
	return squash(n, cur[1:], unicode.IsSpace(r), curTotal+1, total)
}

func squash2(orig []byte) []byte {

	var s = string(orig[:len(orig)])

	var res = ""
	var prevIsSpace = false
	for _, r := range s {

		if unicode.IsSpace(r) && !prevIsSpace {
			prevIsSpace = true
			res = res + " "
		}

		if !unicode.IsSpace(r) {
			prevIsSpace = false
			res = res + string(r)
		}
	}

	return []byte(res)
}

func main() {
	var test = []byte("    This   is the test        data")
	fmt.Println(string(test[:len(test)]))
	var res = squash(test[:1], test[1:], false, 0, len(test))
	fmt.Println(string(res[:len(res)]))
	var res2 = squash2(test)
	fmt.Println(string(res2[:len(res2)]))

}
