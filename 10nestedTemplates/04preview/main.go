package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/favcion.ico", http.NotFoundHandler())

	http.ListenAndServe(":80", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
