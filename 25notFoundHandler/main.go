package main

import (
	"net/http"
	"io"
)

func main() {

	http.Handle("/favcion.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}
func dogImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, "without favicon")
}
