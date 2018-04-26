package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/rivne.png", rivnePngHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	var body string

	body = "<h1>Hello</h1>"
	body += "<br>"
	body += `<img src="/rivne.png">`

	io.WriteString(w, body)
}

func rivnePngHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "../../rivne.png")
}
