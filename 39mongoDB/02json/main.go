package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"fmt"

	"strings"

	"strconv"

	"github.com/slevchyk/goWebDev/39mongoDB/02json/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user/", userHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func userHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	xsPath := strings.Split(path, "/")

	if len(xsPath) < 3 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userIdURL := xsPath[2]

	if userIdURL == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userID, err := strconv.Atoi(userIdURL)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := models.User{
		"James",
		"Male",
		45,
		userID,
	}

	bs, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}

	w.Write(bs)
}
