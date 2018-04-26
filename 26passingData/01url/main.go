package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHendaler)
	http.ListenAndServe(":8080", nil)
}
func indexHendaler(w http.ResponseWriter, r *http.Request) {

	val := r.FormValue("q")
	io.WriteString(w, "my search:"+val)
}

//for test use localhost:8080/?q=itswork
