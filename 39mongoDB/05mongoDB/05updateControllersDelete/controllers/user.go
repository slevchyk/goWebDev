package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/slevchyk/goWebDev/39mongoDB/05mongoDB/05updateControllersDelete/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{
		Session: s,
	}
}

func (uc UserController) UserHandler(w http.ResponseWriter, r *http.Request) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	//store the user in mongo db
	uc.Session.DB("godata").C("users").Insert(u)

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
		getUserHandler(w, r, uc)
	case http.MethodDelete:
		deleteuserHandler(w, r, uc)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}

func getUserHandler(w http.ResponseWriter, r *http.Request, uc UserController) {

	w.Header().Set("Content-Type", "application/json")

	userId, err := getUserId(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !bson.IsObjectIdHex(userId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(userId)

	u := models.User{}

	err = uc.Session.DB("godata").C("users").FindId(oid).One(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("error db query", userId)
		return
	}

	userJson, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}

	w.Write(userJson)
}

func deleteuserHandler(w http.ResponseWriter, r *http.Request, uc UserController) {

	userId, err := getUserId(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !bson.IsObjectIdHex(userId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(userId)

	err = uc.Session.DB("godata").C("users").RemoveId(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintln(w, "Deleted user id", userId)
}

func getUserId(r *http.Request) (string, error) {
	var userId string

	path := r.URL.Path
	xsPath := strings.Split(path, "/")

	if len(xsPath) < 3 {
		return userId, errors.New("Can't get user ID from URL path")
	}

	userId = xsPath[2]

	if userId == "" {
		return userId, errors.New("User ID from URL path is empty")
	}

	return userId, nil
}
