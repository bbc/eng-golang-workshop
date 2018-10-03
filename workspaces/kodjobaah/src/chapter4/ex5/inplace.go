package main

import "fmt"

func inplace(orig []string, cur []string, prev string, curTotal int, total int) []string {

	if curTotal == total {
		return orig
	}

	if cur[0] == prev {
		return inplace(orig, cur[1:], prev, curTotal+1, total)
	}

	n := make([]string, len(orig)+1, cap(orig)+1)
	copy(n, orig)
	n[len(orig)] = cur[0]
	return inplace(n, cur[1:], cur[0], curTotal+1, total)
}

func main() {
	s := [...]string{"0", "1", "1", "3", "4", "4", "5", "5"}
	fmt.Println(s)
	fmt.Println(inplace(s[:1], s[1:], s[0], 0, len(s)-1))
}
