package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	// generate 2 SHA256 hashes
	h1 := sha256.Sum256([]byte("X"))
	h2 := sha256.Sum256([]byte("x"))

	// Count the bits that are different
	var total int
	for i, v := range h1 {
		diff := v ^ h2[i]
		diffPC := pc[diff]
		total += int(diffPC)
		fmt.Printf("%b ^ %b = %b\n", v, h2[i], diff)
		fmt.Printf("different bits: %d\n", diffPC)
		fmt.Printf("total diff: %d\n", total)
	}
	fmt.Println("Total different bits:", total)
}
