package main

import (
	"fmt"
	"net/http"
)

type myHandler bool

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	switch path {
	case "/cow":
		fmt.Fprintf(w, "muuu mu muuuuu")
	case "/frog":
		fmt.Fprintf(w, "rebit rebit")
	default:
		fmt.Fprintf(w, "404")
	}

}

func main() {

	var mh myHandler
	http.ListenAndServe(":8080", mh)
}
