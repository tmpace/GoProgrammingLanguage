// Server4 returns a lissajous gif to the user.
package main

import (
	"net/http"

	"github.com/tmpace/GoProgrammingLanguage/ch1/lissajous"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous.LissaJous(w)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}
