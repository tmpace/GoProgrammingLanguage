package testing

import (
	"fmt"
	"strings"
	"testing"
)

func main() {
	fmt.Println("Testing...")
}

var arr = make([]string, 0)

func tryConcat() {
	var s, sep string

	for i := 0; i < len(arr); i++ {
		s += arr[i] + sep
		sep = " "
	}
}

func tryJoin() {
	strings.Join(arr, " ")
}

func BenchmarkTryConcat(b *testing.B) {
	for i := 0; i < 100000; i++ {
		arr = append(arr, "A")
	}

	for i := 0; i < b.N; i++ {
		tryConcat()
	}
}

func BenchmarkTryJoin(b *testing.B) {
	for i := 0; i < 100000; i++ {
		arr = append(arr, "A")
	}

	for i := 0; i < b.N; i++ {
		tryJoin()
	}
}
