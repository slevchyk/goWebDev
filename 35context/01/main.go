package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHamdler)
	http.ListenAndServe(":8080", nil)
}

func indexHamdler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Fprintln(w, ctx)
}
