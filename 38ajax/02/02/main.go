package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ajax", ajaxhandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func ajaxhandler(w http.ResponseWriter, r *http.Request) {
	s := `Here is some text`
	fmt.Fprintln(w, s)
}
