package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pageToRedirect", pageToRedirectHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying open index page / Method:", r.Method, "\n\n")
}

func pageToRedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying open page for redirect. Method:", r.Method, "\n\n")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
