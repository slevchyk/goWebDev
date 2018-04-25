package main

import (
	"net/http"
	"io"
)

func main() {

	http.Handle("/pub/", http.StripPrefix("/pub", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	var body string

	body = "<h1>Hello</h1>"
	body += "<br>"
	body += `<img src="/pub/rivne.png">`

	io.WriteString(w, body)
}

