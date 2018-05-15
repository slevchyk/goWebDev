package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/slevchyk/goWebDev/39mongoDB/04controllers/models"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) UserHandler(w http.ResponseWriter, r *http.Request) {

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

func (uc UserController) UserHoldsHandler(w http.ResponseWriter, r *http.Request) {

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
