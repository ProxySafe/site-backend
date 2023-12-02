package main

import (
	"fmt"
	"net/http"
)

func main() {
	// entrypoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	http.ListenAndServe(":8003", nil)
}
