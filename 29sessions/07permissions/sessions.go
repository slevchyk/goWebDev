package main

import "net/http"

func getUser(r *http.Request) user {

	var u user

	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	sessionID := c.Value
	userID, ok := dbSessions[sessionID]
	if !ok {
		return u
	}

	u = dbUsers[userID]
	return u
}

func alreadyLoggedIn(r *http.Request) bool {

	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	sessionID := c.Value
	userID := dbSessions[sessionID]
	_, ok := dbUsers[userID]

	return ok
}
