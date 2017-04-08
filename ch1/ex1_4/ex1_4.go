// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}

			countLines(f, counts, fileNames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t\n", n, line)
		}
	}

	for fName, dupes := range fileNames {
		fmt.Printf("File: %s Values: %v\n", fName, dupes)
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		text := input.Text()
		counts[text]++
		fileNames[f.Name()] = append(fileNames[f.Name()], text)
	}
	// NOTE: ignoring potential errors from input.Err()
}
