package main

import (
	"html/template"
	"net/http"

	"fmt"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName  string
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
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		sessionID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, c)
	}

	var u user

	sessionID := c.Value
	userID, ok := dbSessions[sessionID]
	if ok {
		u = dbUsers[userID]
	}

	if r.Method == http.MethodPost {
		userName := r.FormValue("userName")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")

		u = user{
			userName,
			firstName,
			lastName,
		}

		dbSessions[sessionID] = userName
		dbUsers[userName] = u
	}

	fmt.Println("try to execute tpl")
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessionID := c.Value
	userID, ok := dbSessions[sessionID]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	currentUser := dbUsers[userID]
	tpl.ExecuteTemplate(w, "page.gohtml", currentUser)
}
