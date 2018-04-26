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

	http.HandleFunc("/", indexHemdler)
	http.ListenAndServe(":8080", nil)
}

func indexHemdler(w http.ResponseWriter, r *http.Request) {

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	text := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
