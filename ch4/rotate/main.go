package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr = rotate(arr, 13)
	fmt.Println(arr)
}

func rotate(arr []int, shift int) []int {
	// Calculate an offset, the offset
	// is equal to the amount of times
	// our length will go into our shift amount
	offset := shift / len(arr)

	// If the shift amount is higher than the
	// length of the slice, we can calculate a new shift
	// this limits the amount of operations we have to do
	if shift > len(arr) {
		shift = shift - (len(arr) * offset)
	}

	// make a  new slice to hold the result
	result := make([]int, len(arr))

	// loop through the slice and calculate a new position
	for i, v := range arr {
		pos := i + shift

		// account for overflowing the index range
		if pos > len(arr)-1 {
			pos = pos - len(arr)
		}

		result[pos] = v
	}

	// return our newly created result
	return result
}
