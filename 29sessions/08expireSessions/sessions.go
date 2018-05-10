package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, r *http.Request) user {

	var u user

	c, err := r.Cookie("session")
	if err != nil {
		return u
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	sessionID := c.Value
	currentSession, ok := dbSessions[sessionID]
	if ok {
		currentSession.lastActivity = time.Now()
		dbSessions[sessionID] = currentSession
		u = dbUsers[currentSession.userID]
	}

	return u
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {

	if time.Now().Sub(sessionsCleaned) > (time.Duration(sessionLength) * time.Second) {
		go cleanSession()
	}

	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	sessionID := c.Value
	currentSession, ok := dbSessions[sessionID]
	if ok {
		currentSession.lastActivity = time.Now()
		dbSessions[sessionID] = currentSession
	}

	_, ok = dbUsers[currentSession.userID]
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	return ok
}
func cleanSession() {

	fmt.Println("before clean")
	showSessions()

	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Duration(sessionLength) * time.Second) {
			delete(dbSessions, k)
		}
	}
	sessionsCleaned = time.Now()

	fmt.Println("after clean")
	showSessions()

}
func showSessions() {

	for k, v := range dbSessions {
		fmt.Println(k, v.userID)
	}

	fmt.Printf("\n\n")
}
