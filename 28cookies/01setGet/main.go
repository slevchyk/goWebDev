package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/read", readHandler)
	http.ListenAndServe(":8080", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	c := &http.Cookie{
		Name: "my-cookie",
		Value: "some value of my cookie",
	}
	http.SetCookie(w, c)

	io.WriteString(w, "COOKIEE WRITTEN")

}

func readHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	fmt.Fprintln(w, "Your cookie is ", c)

}
