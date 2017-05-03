package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "hello   world    how    r u"
	str = squash([]byte(str))
	fmt.Printf("%q\n", str)
}

func squash(s []byte) string {
	var spaceFound bool
	var spaceStart int
	var numSpaces int

	for i, v := range s {
		// If a space is found in the loop, we check to see if spaceFound is false
		// If it is false, we set it to true, set the start index, and increment numSpaces
		// If it is true, we simply increment numSpaces
		// Either way, if a space was found, we do not want to run the next block so we continue
		// to the next loop iteration
		if unicode.IsSpace(rune(v)) {
			if !spaceFound {
				spaceFound = true
				spaceStart = i
				numSpaces++
			} else {
				numSpaces++
			}

			continue
		}

		// This block will only be run if did NOT find a space in an iteration
		// of the loop AND we have found more than 1 space in a row
		// we then modify the slice to only keep one of the spaces by using copy
		// to slide the higher-numbered elements down by numSpaces to fill the gaps
		// and then we cut off numSpaces+1 elemnts from the end, and recursively call
		// the function
		if spaceFound && numSpaces > 1 {
			fmt.Printf("Found %v spaces starting at index %v squashing.\n", numSpaces, spaceStart)
			copy(s[spaceStart+1:], s[spaceStart+numSpaces:])
			s = s[:len(s)-numSpaces+1]
			fmt.Println("New string: " + string(s))
			return squash(s)
		}

		// If we get here, this means we were at an element that was not
		// a space, and we have not found more than 1 space in a row
		// meaning the previous character was either not a space, or one
		// singular space. In this case, we want to reset the spaceFound
		// and numSpaces variables so that we do not screw up further iterations
		spaceFound = false
		numSpaces = 0
	}

	// finally, if we managed to get through the entire slice
	// without finding a grouping of more than one unicode space,
	// return a string conversion of the byte slice
	return string(s)
}
