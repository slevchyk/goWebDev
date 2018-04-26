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
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	FName := r.FormValue("FName")
	LName := r.FormValue("LName")
	OK := r.FormValue("OK")

	data := struct {
		FName string
		LName string
		OK    string
	}{
		FName,
		LName,
		OK,
	}

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
