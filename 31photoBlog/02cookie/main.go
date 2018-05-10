package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	c := getCookie(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {

		sessionID, _ := uuid.NewV4()

		c = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
	}
	http.SetCookie(w, c)

	return c
}
