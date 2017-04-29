package main

import "fmt"

func main() {
	s := []string{"one", "one", "two", "three", "four", "four", "four"}
	removeDupes(s)
	fmt.Printf("%q\n", s)
}

func removeDupes(s []string) {
	for i := 0; i < len(s)-1; i++ {
		isNext := i < len(s)-2
		isPrev := i > 0
		focus := s[i]

		fmt.Printf("i: %d isNext: %v isPrev: %v\n", i, isNext, isPrev)

		if isNext {
			next := s[i+1]
			if focus == next {
				s[i+1] = ""
			}
		}

		if isPrev {
			prev := s[i-1]
			if focus == prev {
				s[i-1] = ""
			}
		}
	}
}
