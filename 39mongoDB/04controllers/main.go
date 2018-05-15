package main

import (
	"html/template"
	"net/http"

	"github.com/slevchyk/goWebDev/39mongoDB/04controllers/controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))
}

func main() {

	rc := controllers.NewRootController(tpl)
	uc := controllers.NewUserController()

	http.HandleFunc("/", rc.IndexHandler)
	http.HandleFunc("/user", uc.UserHandler)
	http.HandleFunc("/user/", uc.UserHoldsHandler)
	http.ListenAndServe(":8080", nil)
}
