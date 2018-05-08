package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/dog/pics", dogpicsHandler)
	http.HandleFunc("/cat", catHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func dogHandler(w http.ResponseWriter, r *http.Request) {

	c := &http.Cookie{
		Name:  "user-cookie",
		Value: "dog cookie",
		Path:  "/",
	}
	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "dog.gohtml", c)
}

func dogpicsHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
	}

	tpl.ExecuteTemplate(w, "dogpics.gohtml", c)
}

func catHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
	}

	tpl.ExecuteTemplate(w, "cat.gohtml", c)
}
