package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogImg)
	http.ListenAndServe(":8080", nil)

}
func dogImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}

func dog(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", nil)
	if err != nil {
		http.Error(w, "page not found", 404)
	}
}
func foo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, "foo ran")
}
