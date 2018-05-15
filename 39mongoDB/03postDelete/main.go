package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"fmt"

	"strings"

	"strconv"

	"github.com/slevchyk/goWebDev/39mongoDB/03postDelete/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/", userHoldsHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func userHandler(w http.ResponseWriter, r *http.Request) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = 700

	uJson, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	w.Write(uJson)

}

func userHoldsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getUserHandler(w, r)
	case http.MethodDelete:
		deleteuserHandler(w, r)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	xsPath := strings.Split(path, "/")

	if len(xsPath) < 3 {

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

func deleteuserHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintln(w, "Write code to delete user")
}
