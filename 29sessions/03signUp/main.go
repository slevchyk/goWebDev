package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //user id to users
var dbSessions = map[string]string{} //session id to user id

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/page", pageHandler)
	http.HandleFunc("/signup", signupHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	u := getUser(r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "page.gohtml", u)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {

	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {

		userName := r.FormValue("userName")
		password := r.FormValue("password")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")

		if _, ok := dbUsers[userName]; ok {
			http.Error(w, "Usarname already taken", http.StatusForbidden)
			return
		}

		sessionID, _ := uuid.NewV4()

		c := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, c)

		u := user{
			userName,
			password,
			firstName,
			lastName,
		}

		dbSessions[sessionID.String()] = userName
		dbUsers[userName] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
