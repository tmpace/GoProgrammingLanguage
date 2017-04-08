package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for i, arg := range args {
		fmt.Printf("Index: %v Argument: %v\n", i, arg)
	}
}
