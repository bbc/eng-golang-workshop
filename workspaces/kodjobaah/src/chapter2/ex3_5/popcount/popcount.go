// reference used for popcount: https://arxiv.org/pdf/1611.07612.pdf
package popcount

// pc[i] is the population count of 1
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var count byte
	var i uint
	for i = 0; i < 8; i++ {
		count = count + pc[byte(x>>(i*8))]
	}

	return int(count)
}

func PopCountShift(x uint64) int {
	var count uint64 = 0
	var i uint
	for i = 0; i < 63; i++ {
		count = count + ((x >> i) & 1)
	}
	return int(count)
}

func PopCountClearRightmost(x uint64) int {
	var v = 0
	for x != 0 {
		x &= x - 1
		v++
	}
	return v

}
