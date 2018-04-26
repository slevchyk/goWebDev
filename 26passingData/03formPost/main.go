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

	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, `
		<form method="get">
		<input type="text" name="q">
		<input type="submit">
		</form>
		<br>`+val)
}
