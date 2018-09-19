// Non recursive comma insertion
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {

	var buf bytes.Buffer
	l := len(s)
	for i := 0; l > 0; i++ {
		buf.WriteByte(s[l-1])
		if l != 1 && l+1 < len(s) && (i+1)%3 == 0 {
			buf.WriteString(",")
		}
		l--
	}

	var rev = buf.String()
	var revBuf bytes.Buffer

	for l := len(rev); l > 0; l-- {
		revBuf.WriteByte(rev[l-1])
	}

	return revBuf.String()
}

func main() {
	fmt.Println(comma("1234"))
}
