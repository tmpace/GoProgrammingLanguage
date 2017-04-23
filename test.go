// This file will be used to test different examples
package main

import (
	"fmt"
	"time"
)

const (
	noDelay time.Duration = 0
	timeout               = 5 * time.Minute
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
}
