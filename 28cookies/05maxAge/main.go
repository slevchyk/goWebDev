package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/expire", expireHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, `<a href="/set">Set</a>`)
}

func setHandler(w http.ResponseWriter, r *http.Request) {

	c := &http.Cookie{
		Name: "my-cookie",
		Value: "some value",
	}
	http.SetCookie(w, c)

	fmt.Fprintln(w, `<a href="/read">Read</a>`)

}

func readHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintln(w, `<a href="/expire">Delete</a>`)
	fmt.Fprintln(w, c)
}

func expireHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
