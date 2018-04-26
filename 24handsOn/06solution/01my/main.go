package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {

	fs := http.FileServer(http.Dir("public"))

	http.Handle("/pics/", fs)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
