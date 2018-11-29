package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount returns the population count (number of set bits) of x.
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
    var numBitsSet = 0

    for i := 0; i < 8; i++ {
        numBitsSet += int(pc[byte(x>>uint8(i*8))])
    }

    return numBitsSet
}

func PopCountShift(x uint64) int {
    var numBitsSet = 0

    for i := 0; i < 64; i++ {
        if (x>>uint8(i)) & 1 != 0 {
            numBitsSet++
        }
    }

    return numBitsSet
}

func PopCountClear(x uint64) int {
    var numBitsSet = 0

    for ; x != 0; x &= x - 1 {
        numBitsSet++
    }

    return numBitsSet
}
