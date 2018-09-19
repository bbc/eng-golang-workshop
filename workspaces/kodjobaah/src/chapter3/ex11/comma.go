package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {

	floatLoc := strings.Index(s, ".")
	if floatLoc != -1 {
		//		return comma(s[:floatLoc]) + "." + comma(s[floatLoc+1:])
		return comma(s[:floatLoc]) + s[floatLoc:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]

}

func main() {
	fmt.Println(comma("12345.0013"))

}
