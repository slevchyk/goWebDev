package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	cookieTimes, err := r.Cookie("times")
	if err != nil {
		cookieTimes = &http.Cookie{
			Name: "times",
			Value: "0",
		}
	}

	n, err := strconv.Atoi(cookieTimes.Value)
	if err != nil {
		n = 0
	}

	n++

	cookieTimes.Value = strconv.Itoa(n)
	http.SetCookie(w, cookieTimes)

	fmt.Fprintln(w, "visited: ", n, " times")
}
