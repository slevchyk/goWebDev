package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
	Items     []string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/marshal", marshalHandler)
	http.HandleFunc("/encode", encodeHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Index page"))
}

func marshalHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		"James",
		"Bond",
		[]string{"Suit", "Gun"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}

	w.Write(json)

}

func encodeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		"James",
		"Bond",
		[]string{"Suit", "Gun"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
