package main

import "fmt"
import "strings"

func main() {
	fmt.Println(basename("a/b/c.go"))
}

// basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	dot := strings.LastIndex(s, ".")

	s = s[slash+1 : dot]

	return s
}
