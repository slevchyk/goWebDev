package main

import (
	"net/http"
	"io"
	"os"
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

	f, err := os.Open("../../rivne.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}