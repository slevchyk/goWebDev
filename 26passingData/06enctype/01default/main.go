package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHendler)
	http.ListenAndServe(":8080", nil)
}

func indexHendler(w http.ResponseWriter, r *http.Request) {

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	text := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
