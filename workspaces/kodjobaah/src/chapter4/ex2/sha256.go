package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var x = flag.Int("hash", 256, "hash algorithm")
var value = flag.String("value", "value", "item to hash")

func main() {
	flag.Parse()
	switch *x {

	case 384:
		c1 := sha512.Sum384([]byte("x"))
		fmt.Printf("%x\n%T\n", c1, c1)

	case 512:
		c1 := sha512.Sum512([]byte("x"))
		fmt.Printf("%x\n%T\n", c1, c1)

	default:
		c1 := sha256.Sum256([]byte("x"))
		fmt.Printf("%x\n%T\n", c1, c1)

	}

}
