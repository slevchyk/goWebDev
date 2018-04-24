package main

import (
	"fmt"
	"net/http"
)

type myHandlerCow bool

func (mhc myHandlerCow) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "muuu mu muuuuu")
}

type myHandlerFrog bool

func (mhf myHandlerFrog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "rebir rebit")
}

func main() {

	var mhc myHandlerCow
	var mhf myHandlerFrog

	serveMux := http.NewServeMux()
	serveMux.Handle("/cow", mhc)
	serveMux.Handle("/frog", mhf)

	http.ListenAndServe(":8080", serveMux)
}
