package popcount

// pc[i] is the population count of 1
var pc [256]byte

func init() {
	for i := range pc {
		//		fmt.Printf("\n i=%v pc[i]=%v pc[i/2]=%v byte(i&1)=%v", i, pc[i], pc[i/2], byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	// fmt.Printf("hey %v\n", pc)
	// fmt.Printf("0*8 %v\n", pc[byte(x>>(0*8))])
	// fmt.Printf("0*8 %v\n", byte(x>>(1*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(2*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(3*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(4*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(5*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(6*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(7*8)))
	// fmt.Printf("0*8 %v\n", byte(x>>(8*8)))

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
