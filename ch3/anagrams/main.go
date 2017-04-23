package main

import "fmt"

func main() {
	fmt.Println(anagrams("abc", "cba"))
	fmt.Println(anagrams("aabbcc", "ccbbaa"))
	fmt.Println(anagrams("1234", "4322"))
}

// anagrams determins whether two strings are anagrams
func anagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range s1 {
		m[v]++
	}

	for _, v := range s2 {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
