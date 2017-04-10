// Server4 returns a lissajous gif to the user.
package main

import (
	"net/http"

	"strconv"

	"github.com/tmpace/GoProgrammingLanguage/ch1/lissajous"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles := r.Form.Get("cycles")
		if cycles == "" {
			cycles = "5"
		}

		cyclesInt, err := strconv.Atoi(cycles)

		if err != nil {
			panic(err)
		}

		lissajous.LissaJous(w, cyclesInt)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}
