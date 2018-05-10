package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
	Role      string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //user id to users
var dbSessions = map[string]string{} //session id to user id

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))

	enPass, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
	dbUsers["test@domain.com"] = user{
		"test@domain.com",
		enPass,
		"James",
		"Bond",
		"user",
	}
}

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/page", pageHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
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

	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the page", http.StatusForbidden)
		return
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
		role := r.FormValue("role")

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

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Can't encrypt password", http.StatusInternalServerError)
			return
		}

		u := user{
			userName,
			encryptedPassword,
			firstName,
			lastName,
			role,
		}

		dbSessions[sessionID.String()] = userName
		dbUsers[userName] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {

		userName := r.FormValue("userName")
		password := r.FormValue("password")

		//check user
		u, ok := dbUsers[userName]
		if !ok {
			http.Error(w, "Usarname do not patch", http.StatusForbidden)
			return
		}

		//check password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Password do not match", http.StatusForbidden)
			return
		}

		//create session
		sessionID, _ := uuid.NewV4()

		c := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, c)

		dbSessions[sessionID.String()] = userName

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessionID := c.Value
	delete(dbSessions, sessionID)

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
