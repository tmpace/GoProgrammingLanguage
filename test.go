// This file will be used to test different examples
package main

import "fmt"

func main() {
	// declare an array whose length will be set to
	// the number of elements it's initialized with
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}

	Q2 := months[4:7]     // ["April", "May", "June"]
	summer := months[6:9] // ["June", "July", "August"]

	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s is in both Q2 and summer\n", s)
			}
		}
	}

	endlessSummer := summer[:5] // ["June", "July", "August", "September", "October"]
	fmt.Println(endlessSummer)
}
