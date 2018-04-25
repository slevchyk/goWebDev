package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)

}

func meHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my name is Serhii")
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "woof woof")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
