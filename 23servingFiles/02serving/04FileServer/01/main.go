package main

import (
	"net/http"
	"io"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/index", indexHandler)
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

