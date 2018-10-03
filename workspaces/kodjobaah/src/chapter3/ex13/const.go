// Taken from https://github.com/golang/go/wiki/Iota
package main

import "fmt"

type ByteSize float64

const (
	_           = iota // ignore first value which is zero
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	// HMMM
	items := [...]ByteSize{KB, MB, GB, TB, PB, EB, ZB, YB}

	for _, v := range items {
		fmt.Printf("%v\n", v)
		//		fmt.Printf("%b,%v\n", v, v)
	}
}
