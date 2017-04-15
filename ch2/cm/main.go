// Cm takes command-line arguments, or, if none are provided
// it takes arguments from stdin and converts them to their
// values in Feet and Meters
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tmpace/GoProgrammingLanguage/ch2/lenconv"
)

func main() {
	// If no command line args were found
	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a length: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Println(handleConv(text))
	}

	// If we had command line arguments
	for _, str := range os.Args[1:] {
		fmt.Println(handleConv(str))
	}
}

func handleConv(s string) string {
	len, err := strconv.ParseFloat(s, 64)
	if err != nil {
		handleError(err)
	}
	f := lenconv.Feet(len)
	m := lenconv.Meters(len)
	return fmt.Sprintf("%s = %s, %s = %s", f, lenconv.FToM(f), m, lenconv.MToF(m))
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "cm: %v", err)
	os.Exit(1)
}
