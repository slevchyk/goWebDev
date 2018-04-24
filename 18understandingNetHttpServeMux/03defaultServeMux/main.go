package main

import (
	"fmt"
	"net/http"
)

func cowHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "muuu mu muuuuu")
}

func frogHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "rebir rebit")
}

func main() {

	http.HandleFunc("/cow", cowHandler)
	http.HandleFunc("/frog", frogHandler)

	http.ListenAndServe(":8080", nil)
}
