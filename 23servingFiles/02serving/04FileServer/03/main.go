package main

import (
	"io"
	"net/http"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/pub", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	var body string

	body = "<h1>Hello</h1>"
	body += "<br>"
	body += `<img src="/assets/rivne.png">`

	io.WriteString(w, body)
}
