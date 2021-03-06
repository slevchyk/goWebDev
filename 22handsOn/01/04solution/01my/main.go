package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)

}

func meHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "my name is Serhii")
	if err != nil {
		log.Fatalln(err)
	}
}

func dogHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "woof woof")
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "welcome")
	if err != nil {
		log.Fatalln(err)
	}
}
