package main

import (
	"html/template"
	"net/http"

	"strings"

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
	c = appendValue(w, c)
	ss := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", ss)
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

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {

	p1 := "paris.jpg"
	p2 := "london.jpg"
	p3 := "kyiv.jpg"

	s := c.Value

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	c.Value = s
	http.SetCookie(w, c)

	return c
}
