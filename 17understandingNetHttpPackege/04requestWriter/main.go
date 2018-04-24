package main

import (
	"fmt"
	"net/http"
)

type myHandler string

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("My owm header key", "slevchyk")
	//w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Type", "text/plain")

	data := "<h1>Some text to show how plain text looks like</h1>"

	fmt.Fprintf(w, data)
}

func main() {

	var mh myHandler

	http.ListenAndServe(":8080", mh)
}
