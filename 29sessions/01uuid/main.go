package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {

		id, _ := uuid.NewV4()

		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure:true,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}

	fmt.Println(c)
}
