// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	if len(s) < 4 {
		return s
	}

	var buf bytes.Buffer

	idx := len(s) % 3
	buf.WriteString(s[:idx])

	for i := idx; i < len(s); i += 3 {
		buf.WriteString("," + s[i:i+3])
	}

	return buf.String()
}
