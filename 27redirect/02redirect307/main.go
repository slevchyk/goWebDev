package main

import (
	"fmt"
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
	http.HandleFunc("/pageToRedirect", pageToRedirectHandler)
	http.HandleFunc("/pageDestination", pageDestinationHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying open index page / Method:", r.Method, "\n\n")
}

func pageToRedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying open page for redirect. Method:", r.Method, "\n\n")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func pageDestinationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Opening destination page. Method:", r.Method, "\n\n")
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
