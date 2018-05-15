package main

import (
	"html/template"
	"net/http"

	"github.com/slevchyk/goWebDev/39mongoDB/05mongoDB/04updateControllersGet/controllers"
	"gopkg.in/mgo.v2"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))
}

func main() {

	rc := controllers.NewRootController(tpl)
	uc := controllers.NewUserController(getSession())

	http.HandleFunc("/", rc.IndexHandler)
	http.HandleFunc("/user", uc.UserHandler)
	http.HandleFunc("/user/", uc.UserHoldsHandler)
	http.ListenAndServe(":8080", nil)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return s
}
