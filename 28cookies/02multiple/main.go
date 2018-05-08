package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/multiple", multipleHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	c := &http.Cookie{
		Name: "cookie1",
		Value: "some value of my cookie",
	}
	http.SetCookie(w, c)

	io.WriteString(w, "COOKIEE WRITTEN")

}

func readHandler(w http.ResponseWriter, r *http.Request) {

	c1, err := r.Cookie("cookie1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	} else {
		fmt.Fprintln(w, "Your cookie is ", c1)
	}

	c2, err := r.Cookie("cookie2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	} else {
		fmt.Fprintln(w, "Your cookie is ", c2)
	}

	c3, err := r.Cookie("cookie3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	} else {
		fmt.Fprintln(w, "Your cookie is ", c3)
	}
}

func multipleHandler(w http.ResponseWriter, r *http.Request) {

	c2 := &http.Cookie{
		Name: "cookie2",
		Value: "some value of my cookie2",
	}
	http.SetCookie(w, c2)

	c3 := &http.Cookie{
		Name: "cookie3",
		Value: "some value of my cookie3",
	}
	http.SetCookie(w, c3)

	io.WriteString(w, "COOKIEE WRITTEN")
}