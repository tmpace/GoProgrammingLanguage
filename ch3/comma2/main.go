package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-12561782568.12"))
}

// comma adds commas to an number string and ignores decimals and signs
func comma(s string) string {
	// store sign
	var sign string
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = s[:1]
		s = s[1:]
	}

	// store comma and trailing numbers
	var decValues string
	decIndex := strings.LastIndex(s, ".")
	if decIndex > -1 {
		decValues = s[decIndex:]
		s = s[:decIndex]
	}

	var buf bytes.Buffer

	fmt.Println(s)

	for i, v := range s {
		buf.WriteRune(v)
		if i != len(s)-1 && len(s[i+1:])%3 == 0 {
			buf.WriteByte(',')
		}
	}

	return sign + buf.String() + decValues
}

// 1234

// 1. buf = "1", len(234) % 3 = 0, buf = "1,"
