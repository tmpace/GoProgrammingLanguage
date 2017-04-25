package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var shaType = flag.Int("type", 256, "accepts values 384 and 512, defaults to 256")

func main() {
	flag.Parse()
	fmt.Print("Please provide a string to hash: ")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	switch *shaType {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
	default:
		fmt.Printf("%d is not a valid type for a SHA\n", *shaType)
	}
}
