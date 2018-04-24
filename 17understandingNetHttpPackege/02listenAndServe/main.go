package main

import (
	"fmt"
	"net/http"
)

type someType string

func (st someType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " You can do any you want")
}

func main() {

	var st someType

	http.ListenAndServe(":8080", st)
}
