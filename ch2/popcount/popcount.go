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

// PopCountLoop implements the same PopCount method with a loop
func PopCountLoop(x uint64) int {
	var result int
	var i uint

	for ; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

// PopCountShift implements the same PopCount method with bit shifting
func PopCountShift(x uint64) int {
	var i uint
	result := 0

	for ; i < 64; i++ {
		if (x>>i)&1 == 1 {
			result++
		}
	}

	return result
}
