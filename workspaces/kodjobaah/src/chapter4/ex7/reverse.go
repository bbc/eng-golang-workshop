// reverse characters
package main

import "fmt"

func reverse(s []byte) []byte {

	r := []rune(string(s))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return []byte(string(r))

}

func main() {

	var test = []byte("Reverse me")
	fmt.Println(string(test[:len(test)]))
	var res2 = reverse(test)
	fmt.Println(string(res2[:len(res2)]))
}
